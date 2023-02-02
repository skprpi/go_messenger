// Code generated by gowrap. DO NOT EDIT.
// template: ../../../../core/pkg/decorators/templates/logerror_logger.go
// gowrap: http://github.com/hexdigest/gowrap

package decorators

//go:generate gowrap gen -p github.com/YFatMR/go_messenger/auth_service/internal/services -i AccountService -t ../../../../core/pkg/decorators/templates/logerror_logger.go -o logerror_logger.gen.go -l ""

import (
	"context"

	"github.com/YFatMR/go_messenger/auth_service/internal/entities/accountid"
	"github.com/YFatMR/go_messenger/auth_service/internal/entities/credential"
	"github.com/YFatMR/go_messenger/auth_service/internal/entities/token"
	"github.com/YFatMR/go_messenger/auth_service/internal/entities/tokenpayload"
	"github.com/YFatMR/go_messenger/auth_service/internal/services"
	"github.com/YFatMR/go_messenger/core/pkg/errors/logerr"
	"github.com/YFatMR/go_messenger/core/pkg/loggers"
)

// LoggingAccountServiceDecorator implements services.AccountService that is instrumented with custom zap logger
type LoggingAccountServiceDecorator struct {
	logger *loggers.OtelZapLoggerWithTraceID
	base   services.AccountService
}

// NewLoggingAccountServiceDecorator instruments an implementation of the services.AccountService with simple logging
func NewLoggingAccountServiceDecorator(base services.AccountService, logger *loggers.OtelZapLoggerWithTraceID) *LoggingAccountServiceDecorator {
	if base == nil {
		panic("LoggingAccountServiceDecorator got empty base")
	}
	if logger == nil {
		panic("LoggingAccountServiceDecorator got empty logger")
	}
	return &LoggingAccountServiceDecorator{
		base:   base,
		logger: logger,
	}
}

// CreateAccount implements services.AccountService
func (d *LoggingAccountServiceDecorator) CreateAccount(ctx context.Context, credential *credential.Entity) (accountID *accountid.Entity, lerr logerr.Error) {

	d.logger.DebugContextNoExport(ctx, "LoggingAccountServiceDecorator: calling CreateAccount")
	defer func() {
		defer d.logger.DebugContextNoExport(ctx, "LoggingAccountServiceDecorator: CreateAccount finished")
		if lerr == nil {
			return
		}
		if lerr.IsLogMessage() {
			d.logger.LogContextLogerror(ctx, lerr)
			lerr.StopLogMessage()
		}
	}()
	return d.base.CreateAccount(ctx, credential)
}

// GetToken implements services.AccountService
func (d *LoggingAccountServiceDecorator) GetToken(ctx context.Context, credential *credential.Entity) (token *token.Entity, lerr logerr.Error) {

	d.logger.DebugContextNoExport(ctx, "LoggingAccountServiceDecorator: calling GetToken")
	defer func() {
		defer d.logger.DebugContextNoExport(ctx, "LoggingAccountServiceDecorator: GetToken finished")
		if lerr == nil {
			return
		}
		if lerr.IsLogMessage() {
			d.logger.LogContextLogerror(ctx, lerr)
			lerr.StopLogMessage()
		}
	}()
	return d.base.GetToken(ctx, credential)
}

// GetTokenPayload implements services.AccountService
func (d *LoggingAccountServiceDecorator) GetTokenPayload(ctx context.Context, token *token.Entity) (tokenPayload *tokenpayload.Entity, lerr logerr.Error) {

	d.logger.DebugContextNoExport(ctx, "LoggingAccountServiceDecorator: calling GetTokenPayload")
	defer func() {
		defer d.logger.DebugContextNoExport(ctx, "LoggingAccountServiceDecorator: GetTokenPayload finished")
		if lerr == nil {
			return
		}
		if lerr.IsLogMessage() {
			d.logger.LogContextLogerror(ctx, lerr)
			lerr.StopLogMessage()
		}
	}()
	return d.base.GetTokenPayload(ctx, token)
}