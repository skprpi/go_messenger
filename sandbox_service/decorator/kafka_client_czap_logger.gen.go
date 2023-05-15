// Code generated by gowrap. DO NOT EDIT.
// template: ../../core/pkg/decorators/templates/czap_logger.template.go
// gowrap: http://github.com/hexdigest/gowrap

package decorator

//go:generate gowrap gen -p github.com/YFatMR/go_messenger/sandbox_service/apientity -i KafkaClient -t ../../core/pkg/decorators/templates/czap_logger.template.go -o kafka_client_czap_logger.gen.go -l ""

import (
	"context"

	"github.com/YFatMR/go_messenger/core/pkg/czap"
	"github.com/YFatMR/go_messenger/sandbox_service/apientity"
	"github.com/YFatMR/go_messenger/sandbox_service/entity"
	"go.uber.org/zap"
)

// LoggingKafkaClientDecorator implements apientity.KafkaClient that is instrumented with custom zap logger
type LoggingKafkaClientDecorator struct {
	logger *czap.Logger
	base   apientity.KafkaClient
}

// NewLoggingKafkaClientDecorator instruments an implementation of the apientity.KafkaClient with simple logging
func NewLoggingKafkaClientDecorator(base apientity.KafkaClient, logger *czap.Logger) *LoggingKafkaClientDecorator {
	if base == nil {
		panic("LoggingKafkaClientDecorator got empty base")
	}
	if logger == nil {
		panic("LoggingKafkaClientDecorator got empty logger")
	}
	return &LoggingKafkaClientDecorator{
		base:   base,
		logger: logger,
	}
}

// Stop implements apientity.KafkaClient
func (d *LoggingKafkaClientDecorator) Stop() {

	d.logger.Info("LoggingKafkaClientDecorator: calling Stop")
	defer d.logger.Info("LoggingKafkaClientDecorator: Stop finished")
	d.base.Stop()
	return
}

// WriteCodeRunnerMessage implements apientity.KafkaClient
func (d *LoggingKafkaClientDecorator) WriteCodeRunnerMessage(ctx context.Context, userID *entity.UserID, programID *entity.ProgramID, sourceCode string, language entity.Languages) (err error) {

	d.logger.InfoContext(ctx, "LoggingKafkaClientDecorator: calling WriteCodeRunnerMessage")
	defer func() {
		if err != nil {
			d.logger.ErrorContext(ctx, "", zap.NamedError("public api error", err))
		}
		d.logger.InfoContext(ctx, "LoggingKafkaClientDecorator: WriteCodeRunnerMessage finished")
	}()
	return d.base.WriteCodeRunnerMessage(ctx, userID, programID, sourceCode, language)
}
