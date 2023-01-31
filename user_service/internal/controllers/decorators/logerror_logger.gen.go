// Code generated by gowrap. DO NOT EDIT.
// template: ../../../../core/pkg/decorators/templates/logerror_logger.go
// gowrap: http://github.com/hexdigest/gowrap

package decorators

//go:generate gowrap gen -p github.com/YFatMR/go_messenger/user_service/internal/controllers -i UserController -t ../../../../core/pkg/decorators/templates/logerror_logger.go -o logerror_logger.gen.go -l ""

import (
	"context"

	"github.com/YFatMR/go_messenger/core/pkg/errors/logerr"
	"github.com/YFatMR/go_messenger/core/pkg/loggers"
	"github.com/YFatMR/go_messenger/protocol/pkg/proto"
	"github.com/YFatMR/go_messenger/user_service/internal/controllers"
)

// LoggingUserControllerDecorator implements controllers.UserController that is instrumented with custom zap logger
type LoggingUserControllerDecorator struct {
	logger *loggers.OtelZapLoggerWithTraceID
	base   controllers.UserController
}

// NewLoggingUserControllerDecorator instruments an implementation of the controllers.UserController with simple logging
func NewLoggingUserControllerDecorator(base controllers.UserController, logger *loggers.OtelZapLoggerWithTraceID) *LoggingUserControllerDecorator {
	if base == nil {
		panic("LoggingUserControllerDecorator got empty base")
	}
	if logger == nil {
		panic("LoggingUserControllerDecorator got empty logger")
	}
	return &LoggingUserControllerDecorator{
		base:   base,
		logger: logger,
	}
}

// Create implements controllers.UserController
func (d *LoggingUserControllerDecorator) Create(ctx context.Context, request *proto.CreateUserDataRequest) (userID *proto.UserID, lerr logerr.Error) {

	d.logger.DebugContextNoExport(ctx, "LoggingUserControllerDecorator: calling Create")
	defer func() {
		defer d.logger.DebugContextNoExport(ctx, "LoggingUserControllerDecorator: Create finished")
		if lerr == nil {
			return
		}
		if lerr.IsLogMessage() {
			// TODO: create special template for nil error logick
			// If we have no error, make error nil to prevent logging the same message many times
			d.logger.LogContextLogerror(ctx, lerr)
			lerr.StopLogMessage()
		}
		if !lerr.HasError() {
			lerr = nil
		}
	}()
	return d.base.Create(ctx, request)
}

// DeleteByID implements controllers.UserController
func (d *LoggingUserControllerDecorator) DeleteByID(ctx context.Context, request *proto.UserID) (void *proto.Void, lerr logerr.Error) {

	d.logger.DebugContextNoExport(ctx, "LoggingUserControllerDecorator: calling DeleteByID")
	defer func() {
		defer d.logger.DebugContextNoExport(ctx, "LoggingUserControllerDecorator: DeleteByID finished")
		if lerr == nil {
			return
		}
		if lerr.IsLogMessage() {
			// TODO: create special template for nil error logick
			// If we have no error, make error nil to prevent logging the same message many times
			d.logger.LogContextLogerror(ctx, lerr)
			lerr.StopLogMessage()
		}
		if !lerr.HasError() {
			lerr = nil
		}
	}()
	return d.base.DeleteByID(ctx, request)
}

// GetByID implements controllers.UserController
func (d *LoggingUserControllerDecorator) GetByID(ctx context.Context, request *proto.UserID) (userData *proto.UserData, lerr logerr.Error) {

	d.logger.DebugContextNoExport(ctx, "LoggingUserControllerDecorator: calling GetByID")
	defer func() {
		defer d.logger.DebugContextNoExport(ctx, "LoggingUserControllerDecorator: GetByID finished")
		if lerr == nil {
			return
		}
		if lerr.IsLogMessage() {
			// TODO: create special template for nil error logick
			// If we have no error, make error nil to prevent logging the same message many times
			d.logger.LogContextLogerror(ctx, lerr)
			lerr.StopLogMessage()
		}
		if !lerr.HasError() {
			lerr = nil
		}
	}()
	return d.base.GetByID(ctx, request)
}

// Ping implements controllers.UserController
func (d *LoggingUserControllerDecorator) Ping(ctx context.Context, request *proto.Void) (pong *proto.Pong, lerr logerr.Error) {

	d.logger.DebugContextNoExport(ctx, "LoggingUserControllerDecorator: calling Ping")
	defer func() {
		defer d.logger.DebugContextNoExport(ctx, "LoggingUserControllerDecorator: Ping finished")
		if lerr == nil {
			return
		}
		if lerr.IsLogMessage() {
			// TODO: create special template for nil error logick
			// If we have no error, make error nil to prevent logging the same message many times
			d.logger.LogContextLogerror(ctx, lerr)
			lerr.StopLogMessage()
		}
		if !lerr.HasError() {
			lerr = nil
		}
	}()
	return d.base.Ping(ctx, request)
}
