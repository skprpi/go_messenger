//go:build test
// +build test

package test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/YFatMR/go_messenger/core/pkg/ckafka"
	"github.com/YFatMR/go_messenger/core/pkg/configs/cviper"
	"github.com/YFatMR/go_messenger/protocol/pkg/proto"
	"github.com/stretchr/testify/suite"
)

type SandboxTestSuite struct {
	sandboxTopicKafkaClient KafkaClient
	userManager             UserManager
	suite.Suite
}

func TestSandboxTestSuite(t *testing.T) {
	config := cviper.New()
	config.SetConfigFile(envFile)
	if err := config.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		panic(err)
	}

	sandboxTopicKafkaClient := NewKafkaClientFromConfig(config)
	defer sandboxTopicKafkaClient.Close()

	suite.Run(
		t,
		&SandboxTestSuite{sandboxTopicKafkaClient: sandboxTopicKafkaClient},
	)
}

func (s *SandboxTestSuite) TestCreateProgram() {
	ctx := context.Background()
	require := s.Require()

	_, token, err := s.userManager.NewAuthorizedUser(ctx)
	require.NoError(err)

	program := NewHelloWorldProgram()

	ctx = s.userManager.NewContextWithToken(ctx, token)
	programID, err := frontServicegRPCClient.CreateProgram(ctx, &proto.ProgramSource{
		Language:   "go",
		SourceCode: program.SourceCode,
	})
	require.NoError(err)
	require.NotNil(programID)
}

func (s *SandboxTestSuite) TestGetProgramAfterCreation() {
	ctx := context.Background()
	require := s.Require()

	_, token, err := s.userManager.NewAuthorizedUser(ctx)
	require.NoError(err)

	program := NewHelloWorldProgram()

	ctx = s.userManager.NewContextWithToken(ctx, token)
	programID, err := frontServicegRPCClient.CreateProgram(ctx, &proto.ProgramSource{
		Language:   "go",
		SourceCode: program.SourceCode,
	})
	require.NoError(err)
	require.NotNil(programID)

	programSourceResponse, err := frontServicegRPCClient.GetProgramByID(ctx, programID)
	require.NoError(err)
	require.NotNil(programSourceResponse)

	require.Equal(program.SourceCode, programSourceResponse.Source.SourceCode)
}

func (s *SandboxTestSuite) TestRunHelloWorld() {
	ctx := context.Background()
	require := s.Require()

	_, token, err := s.userManager.NewAuthorizedUser(ctx)
	require.NoError(err)

	expectedProgram := NewHelloWorldProgram()

	ctx = s.userManager.NewContextWithToken(ctx, token)
	programID, err := frontServicegRPCClient.CreateProgram(ctx, &proto.ProgramSource{
		Language:   "go",
		SourceCode: expectedProgram.SourceCode,
	})
	require.NoError(err)
	require.NotNil(programID)

	_, err = frontServicegRPCClient.RunProgram(ctx, programID)
	require.NoError(err)

	// Wait kafka
	message, err := s.sandboxTopicKafkaClient.WaitMessage(ctx)
	require.NoError(err)

	var programExecutionMessage ckafka.ProgramExecutionMessage
	err = json.Unmarshal(message.Value, &programExecutionMessage)
	require.NoError(err)

	// Message got. It means that program is executed and result writed to database
	program, err := frontServicegRPCClient.GetProgramByID(ctx, programID)
	require.NoError(err)
	require.Equal(expectedProgram.Stdout, program.GetCodeRunnerOutput().GetStdout())
	require.Equal(expectedProgram.Stderr, program.GetCodeRunnerOutput().GetStderr())
	require.Equal(expectedProgram.SourceCode, program.GetSource().GetSourceCode())
}
