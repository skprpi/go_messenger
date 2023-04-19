// Code generated by gowrap. DO NOT EDIT.
// template: ../../core/pkg/decorators/templates/czap_logger.template.go
// gowrap: http://github.com/hexdigest/gowrap

package decorator

//go:generate gowrap gen -p github.com/YFatMR/go_messenger/dialog_service/apientity -i DialogController -t ../../core/pkg/decorators/templates/czap_logger.template.go -o dialog_controller_czap_logger.gen.go -l ""

import (
	"context"

	"github.com/YFatMR/go_messenger/core/pkg/czap"
	"github.com/YFatMR/go_messenger/dialog_service/apientity"
	"github.com/YFatMR/go_messenger/protocol/pkg/proto"
	"go.uber.org/zap"
)

// LoggingDialogControllerDecorator implements apientity.DialogController that is instrumented with custom zap logger
type LoggingDialogControllerDecorator struct {
	logger *czap.Logger
	base   apientity.DialogController
}

// NewLoggingDialogControllerDecorator instruments an implementation of the apientity.DialogController with simple logging
func NewLoggingDialogControllerDecorator(base apientity.DialogController, logger *czap.Logger) *LoggingDialogControllerDecorator {
	if base == nil {
		panic("LoggingDialogControllerDecorator got empty base")
	}
	if logger == nil {
		panic("LoggingDialogControllerDecorator got empty logger")
	}
	return &LoggingDialogControllerDecorator{
		base:   base,
		logger: logger,
	}
}

// CreateDialogWith implements apientity.DialogController
func (d *LoggingDialogControllerDecorator) CreateDialogWith(ctx context.Context, request *proto.UserID) (dp1 *proto.Dialog, err error) {

	d.logger.InfoContext(ctx, "LoggingDialogControllerDecorator: calling CreateDialogWith")
	defer func() {
		if err != nil {
			d.logger.ErrorContext(ctx, "", zap.NamedError("public api error", err))
		}
		d.logger.InfoContext(ctx, "LoggingDialogControllerDecorator: CreateDialogWith finished")
	}()
	return d.base.CreateDialogWith(ctx, request)
}

// GetDialogs implements apientity.DialogController
func (d *LoggingDialogControllerDecorator) GetDialogs(ctx context.Context, request *proto.GetDialogsRequest) (response *proto.GetDialogsResponse, err error) {

	d.logger.InfoContext(ctx, "LoggingDialogControllerDecorator: calling GetDialogs")
	defer func() {
		if err != nil {
			d.logger.ErrorContext(ctx, "", zap.NamedError("public api error", err))
		}
		d.logger.InfoContext(ctx, "LoggingDialogControllerDecorator: GetDialogs finished")
	}()
	return d.base.GetDialogs(ctx, request)
}

// Ping implements apientity.DialogController
func (d *LoggingDialogControllerDecorator) Ping(ctx context.Context, request *proto.Void) (pong *proto.Pong, err error) {

	d.logger.InfoContext(ctx, "LoggingDialogControllerDecorator: calling Ping")
	defer func() {
		if err != nil {
			d.logger.ErrorContext(ctx, "", zap.NamedError("public api error", err))
		}
		d.logger.InfoContext(ctx, "LoggingDialogControllerDecorator: Ping finished")
	}()
	return d.base.Ping(ctx, request)
}
