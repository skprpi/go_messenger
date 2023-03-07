// Code generated by gowrap. DO NOT EDIT.
// template: ../../core/pkg/decorators/templates/opentelemetry_tracing.template.go
// gowrap: http://github.com/hexdigest/gowrap

package decorator

//go:generate gowrap gen -p github.com/YFatMR/go_messenger/user_service/apientity -i UserService -t ../../core/pkg/decorators/templates/opentelemetry_tracing.template.go -o user_service_opentelemetry_tracing.gen.go -l ""

import (
	"context"

	"github.com/YFatMR/go_messenger/user_service/apientity"
	"github.com/YFatMR/go_messenger/user_service/entity"
	"go.opentelemetry.io/otel/trace"
)

// OpentelemetryTracingUserServiceDecorator implements apientity.UserService that is instrumented with custom zap logger
type OpentelemetryTracingUserServiceDecorator struct {
	base         apientity.UserService
	tracer       trace.Tracer
	recordErrors bool
}

// NewOpentelemetryTracingUserServiceDecorator instruments an implementation of the apientity.UserService with simple logging
func NewOpentelemetryTracingUserServiceDecorator(base apientity.UserService, tracer trace.Tracer, recordErrors bool) *OpentelemetryTracingUserServiceDecorator {
	if base == nil {
		panic("OpentelemetryTracingUserServiceDecorator got empty base")
	}
	if tracer == nil {
		panic("OpentelemetryTracingUserServiceDecorator got empty tracer")
	}
	return &OpentelemetryTracingUserServiceDecorator{
		base:         base,
		tracer:       tracer,
		recordErrors: recordErrors,
	}
}

// Create implements apientity.UserService
func (d *OpentelemetryTracingUserServiceDecorator) Create(ctx context.Context, user *entity.User, unsafeCredential *entity.UnsafeCredential) (userID *entity.UserID, err error) {

	var span trace.Span
	ctx, span = d.tracer.Start(ctx, "/Create")
	defer func() {
		if err != nil && d.recordErrors {
			span.RecordError(err)
		}
		span.End()
	}()
	return d.base.Create(ctx, user, unsafeCredential)
}

// DeleteByID implements apientity.UserService
func (d *OpentelemetryTracingUserServiceDecorator) DeleteByID(ctx context.Context, userID *entity.UserID) (err error) {

	var span trace.Span
	ctx, span = d.tracer.Start(ctx, "/DeleteByID")
	defer func() {
		if err != nil && d.recordErrors {
			span.RecordError(err)
		}
		span.End()
	}()
	return d.base.DeleteByID(ctx, userID)
}

// GenerateToken implements apientity.UserService
func (d *OpentelemetryTracingUserServiceDecorator) GenerateToken(ctx context.Context, unsafeCredential *entity.UnsafeCredential) (token *entity.Token, err error) {

	var span trace.Span
	ctx, span = d.tracer.Start(ctx, "/GenerateToken")
	defer func() {
		if err != nil && d.recordErrors {
			span.RecordError(err)
		}
		span.End()
	}()
	return d.base.GenerateToken(ctx, unsafeCredential)
}

// GetByID implements apientity.UserService
func (d *OpentelemetryTracingUserServiceDecorator) GetByID(ctx context.Context, userID *entity.UserID) (user *entity.User, err error) {

	var span trace.Span
	ctx, span = d.tracer.Start(ctx, "/GetByID")
	defer func() {
		if err != nil && d.recordErrors {
			span.RecordError(err)
		}
		span.End()
	}()
	return d.base.GetByID(ctx, userID)
}
