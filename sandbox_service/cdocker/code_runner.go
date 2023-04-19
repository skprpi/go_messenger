package cdocker

import (
	"archive/tar"
	"bytes"
	"context"
	"io"
	"strconv"
	"time"

	"github.com/YFatMR/go_messenger/core/pkg/configs/cviper"
	"github.com/YFatMR/go_messenger/core/pkg/czap"
	"github.com/YFatMR/go_messenger/sandbox_service/apientity"
	"github.com/YFatMR/go_messenger/sandbox_service/entity"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"go.uber.org/zap"
)

type DockerRunnerLimitations struct {
	InputSourceCodeByte int
	OutputStdoutByte    int
	OutputStderrByte    int
}

type codeRunner struct {
	client      *client.Client
	config      *DockerRunnerSettings
	limitations *DockerRunnerLimitations
	logger      *czap.Logger
}

type DockerRunnerSettings struct {
	ImageName                   string
	ImagePath                   string
	MemoryLimitBytes            int64
	NetworkDisable              bool
	ContainerNamePrefix         string
	ContainerGoSourcesDirectory string
	RunCodeTimeout              time.Duration
}

func DockerRunnerSettingsFromConfig(config *cviper.CustomViper) *DockerRunnerSettings {
	return &DockerRunnerSettings{
		ImageName:                   config.GetStringRequired("DOCKER_CODE_RUNNER_IMAGE_NAME"),
		ImagePath:                   config.GetStringRequired("DOCKER_CODE_RUNNER_IMAGE_PATH"),
		MemoryLimitBytes:            config.GetInt64Required("DOCKER_CODE_RUNNER_MEMORY_LIMITATION_BYTES"),
		NetworkDisable:              config.GetBoolRequired("DOCKER_CODE_RUNNER_NETWORK_DISABLED"),
		ContainerNamePrefix:         config.GetStringRequired("DOCKER_CODE_RUNNER_CONTAINER_NAME_PREFIX"),
		ContainerGoSourcesDirectory: config.GetStringRequired("DOCKER_CODE_RUNNER_GO_IMAGE_SOURCE_DIRECTORY"),
		RunCodeTimeout: config.GetMillisecondsDurationRequired(
			"DOCKER_CODE_RUNNER_PROGRAM_EXECUTION_TIMEOUT_MILLISECONDS",
		),
	}
}

func DockerRunnerLimitationsFromConfig(config *cviper.CustomViper) *DockerRunnerLimitations {
	return &DockerRunnerLimitations{
		InputSourceCodeByte: config.GetIntRequired("DOCKER_CODE_RUNNER_INPUT_SOURCE_CODE_LIMITATION_BYTE"),
		OutputStdoutByte:    config.GetIntRequired("DOCKER_CODE_RUNNER_OUTPUT_STDOUT_LIMITATION_BYTE"),
		OutputStderrByte:    config.GetIntRequired("DOCKER_CODE_RUNNER_OUTPUT_STDERR_LIMITATION_BYTE"),
	}
}

func NewCodeRunner(client *client.Client, config *DockerRunnerSettings, limitations *DockerRunnerLimitations,
	logger *czap.Logger,
) apientity.CodeRunner {
	return &codeRunner{
		client:      client,
		config:      config,
		limitations: limitations,
		logger:      logger,
	}
}

func CodeRunnerFromConfig(ctx context.Context, config *cviper.CustomViper, logger *czap.Logger) (
	apientity.CodeRunner, error,
) {
	dockerRunnerSettings := DockerRunnerSettingsFromConfig(config)
	dockerRunnerLimitations := DockerRunnerLimitationsFromConfig(config)
	dockerClient, err := ClientFromConfig(config)
	if err != nil {
		return nil, err
	}
	LoadDockerImageFromTAR(ctx, dockerRunnerSettings.ImagePath, dockerClient)
	// It check that the desired docker image has been found
	_, _, err = dockerClient.ImageInspectWithRaw(ctx, dockerRunnerSettings.ImageName)
	if err != nil {
		return nil, err
	}
	return NewCodeRunner(dockerClient, dockerRunnerSettings, dockerRunnerLimitations, logger), nil
}

func (c *codeRunner) Stop() {
	c.client.Close()
}

func (c *codeRunner) RunGoCode(ctx context.Context, sourceCode string, userID *entity.UserID) (
	*entity.ProgramOutput, error,
) {
	if len(sourceCode) > c.limitations.InputSourceCodeByte {
		c.logger.ErrorContext(
			ctx, "Too huge input",
			zap.Int("input code size byte", len(sourceCode)),
			zap.Int("limitation byte", c.limitations.InputSourceCodeByte),
		)
		return nil, ErrHugeInput
	}
	containerName := c.config.ContainerNamePrefix + strconv.FormatUint(userID.ID, 10)
	exeContainer, err := c.createContainer(ctx, containerName)
	if err != nil {
		c.logger.ErrorContext(ctx, "Can not create container", zap.Error(err))
		return nil, ErrRunCode
	}
	containerID := exeContainer.ID
	defer c.client.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{Force: true})

	reader, err := c.createTARReader(sourceCode)
	if err != nil {
		c.logger.ErrorContext(ctx, "Can not create TAR reader", zap.Error(err))
		return nil, ErrRunCode
	}

	ctx, cancel := context.WithTimeout(ctx, c.config.RunCodeTimeout)
	defer cancel()

	err = c.client.CopyToContainer(
		ctx, containerID, c.config.ContainerGoSourcesDirectory, reader,
		types.CopyToContainerOptions{AllowOverwriteDirWithFile: false},
	)
	if err != nil {
		c.logger.ErrorContext(ctx, "Can not copy to container", zap.Error(err))
		return nil, ErrRunCode
	}

	if err = c.client.ContainerStart(ctx, containerID, types.ContainerStartOptions{}); err != nil {
		c.logger.ErrorContext(ctx, "Can not start container", zap.Error(err))
		return nil, ErrRunCode
	}

	waitResponce, errCh := c.client.ContainerWait(ctx, containerID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			c.logger.ErrorContext(ctx, "Gor error from chan", zap.Error(err))
			return nil, ErrRunCode
		}
	case <-waitResponce:
	}

	return c.GetProgramOutput(ctx, containerID)
}

func (c *codeRunner) createTARReader(sourceCode string) (
	io.Reader, error,
) {
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)

	hdr := &tar.Header{
		Name: "main.go",
		Mode: 0o644,
		Size: int64(len(sourceCode)),
	}
	if err := tw.WriteHeader(hdr); err != nil {
		return nil, err
	}
	if _, err := tw.Write([]byte(sourceCode)); err != nil {
		return nil, err
	}

	// do not use defer (it write footer)
	tw.Close()
	return bytes.NewReader(buf.Bytes()), nil
}

func (c *codeRunner) createContainer(ctx context.Context, containerName string) (
	container.CreateResponse, error,
) {
	return c.client.ContainerCreate(
		ctx,
		&container.Config{
			Image:           c.config.ImageName,
			AttachStdout:    true,
			AttachStderr:    true,
			NetworkDisabled: c.config.NetworkDisable,
			// Tty:             false, # ????
			// Cmd:             []string{"ls", "/app"},
		},
		&container.HostConfig{
			Resources: container.Resources{
				Memory: c.config.MemoryLimitBytes,
				// CPUQuota: int64(50 * 1024 * 1024 * 1024),
			},
		},
		nil, nil,
		containerName,
	)
}

func (c *codeRunner) GetProgramOutput(ctx context.Context, containerID string) (
	*entity.ProgramOutput, error,
) {
	logs, err := c.client.ContainerLogs(
		ctx, containerID,
		types.ContainerLogsOptions{
			ShowStdout: true,
			ShowStderr: true,
		},
	)
	if err != nil {
		return nil, err
	}
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	_, err = stdcopy.StdCopy(stdout, stderr, logs)
	if err != nil {
		return nil, err
	}

	limitedStdout := stdout.Next(min(c.limitations.OutputStdoutByte, stdout.Len()))
	limitedStderr := stderr.Next(min(c.limitations.OutputStderrByte, stderr.Len()))

	return &entity.ProgramOutput{
		Stdout: string(limitedStdout),
		Stderr: string(limitedStderr),
	}, nil
}
