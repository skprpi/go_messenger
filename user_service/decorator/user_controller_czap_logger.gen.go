// Code generated by gowrap. DO NOT EDIT.
// template: ../../core/pkg/decorators/templates/czap_logger.template.go
// gowrap: http://github.com/hexdigest/gowrap

package decorator

//go:generate gowrap gen -p github.com/YFatMR/go_messenger/user_service/apientity -i UserController -t ../../core/pkg/decorators/templates/czap_logger.template.go -o user_controller_czap_logger.gen.go -l ""

import (
	"context"

	"github.com/YFatMR/go_messenger/core/pkg/czap"
	"github.com/YFatMR/go_messenger/protocol/pkg/proto"
	"github.com/YFatMR/go_messenger/user_service/apientity"
	"go.uber.org/zap"
)

// LoggingUserControllerDecorator implements apientity.UserController that is instrumented with custom zap logger
type LoggingUserControllerDecorator struct {
	logger *czap.Logger
	base   apientity.UserController
}

// NewLoggingUserControllerDecorator instruments an implementation of the apientity.UserController with simple logging
func NewLoggingUserControllerDecorator(base apientity.UserController, logger *czap.Logger) *LoggingUserControllerDecorator {
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

// Create implements apientity.UserController
func (d *LoggingUserControllerDecorator) Create(ctx context.Context, request *proto.CreateUserRequest) (userID *proto.UserID, err error) {

	d.logger.InfoContext(ctx, "LoggingUserControllerDecorator: calling Create")
	defer func() {
		if err != nil {
			d.logger.ErrorContext(ctx, "", zap.NamedError("public api error", err))
		}
		d.logger.InfoContext(ctx, "LoggingUserControllerDecorator: Create finished")
	}()
	return d.base.Create(ctx, request)
}

// DeleteByID implements apientity.UserController
func (d *LoggingUserControllerDecorator) DeleteByID(ctx context.Context, request *proto.UserID) (void *proto.Void, err error) {

	d.logger.InfoContext(ctx, "LoggingUserControllerDecorator: calling DeleteByID")
	defer func() {
		if err != nil {
			d.logger.ErrorContext(ctx, "", zap.NamedError("public api error", err))
		}
		d.logger.InfoContext(ctx, "LoggingUserControllerDecorator: DeleteByID finished")
	}()
	return d.base.DeleteByID(ctx, request)
}

// GenerateToken implements apientity.UserController
func (d *LoggingUserControllerDecorator) GenerateToken(ctx context.Context, request *proto.Credential) (void *proto.Token, err error) {

	d.logger.InfoContext(ctx, "LoggingUserControllerDecorator: calling GenerateToken")
	defer func() {
		if err != nil {
			d.logger.ErrorContext(ctx, "", zap.NamedError("public api error", err))
		}
		d.logger.InfoContext(ctx, "LoggingUserControllerDecorator: GenerateToken finished")
	}()
	return d.base.GenerateToken(ctx, request)
}

// GetByID implements apientity.UserController
func (d *LoggingUserControllerDecorator) GetByID(ctx context.Context, request *proto.UserID) (userData *proto.UserData, err error) {

	d.logger.InfoContext(ctx, "LoggingUserControllerDecorator: calling GetByID")
	defer func() {
		if err != nil {
			d.logger.ErrorContext(ctx, "", zap.NamedError("public api error", err))
		}
		d.logger.InfoContext(ctx, "LoggingUserControllerDecorator: GetByID finished")
	}()
	return d.base.GetByID(ctx, request)
}

// Ping implements apientity.UserController
func (d *LoggingUserControllerDecorator) Ping(ctx context.Context, request *proto.Void) (pong *proto.Pong, err error) {

	d.logger.InfoContext(ctx, "LoggingUserControllerDecorator: calling Ping")
	defer func() {
		if err != nil {
			d.logger.ErrorContext(ctx, "", zap.NamedError("public api error", err))
		}
		d.logger.InfoContext(ctx, "LoggingUserControllerDecorator: Ping finished")
	}()
	return d.base.Ping(ctx, request)
}
