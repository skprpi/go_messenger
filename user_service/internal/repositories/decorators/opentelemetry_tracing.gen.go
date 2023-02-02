// Code generated by gowrap. DO NOT EDIT.
// template: ../../../../core/pkg/decorators/templates/opentelemetry_tracing.go
// gowrap: http://github.com/hexdigest/gowrap

package decorators

//go:generate gowrap gen -p github.com/YFatMR/go_messenger/user_service/internal/repositories -i UserRepository -t ../../../../core/pkg/decorators/templates/opentelemetry_tracing.go -o opentelemetry_tracing.gen.go -l ""

import (
	"context"

	"github.com/YFatMR/go_messenger/core/pkg/ulo"
	"github.com/YFatMR/go_messenger/user_service/internal/entities/accountid"
	"github.com/YFatMR/go_messenger/user_service/internal/entities/user"
	"github.com/YFatMR/go_messenger/user_service/internal/entities/userid"
	"github.com/YFatMR/go_messenger/user_service/internal/repositories"
	"go.opentelemetry.io/otel/trace"
)

// OpentelemetryTracingUserRepositoryDecorator implements repositories.UserRepository that is instrumented with custom zap logger
type OpentelemetryTracingUserRepositoryDecorator struct {
	base         repositories.UserRepository
	tracer       trace.Tracer
	recordErrors bool
}

// NewOpentelemetryTracingUserRepositoryDecorator instruments an implementation of the repositories.UserRepository with simple logging
func NewOpentelemetryTracingUserRepositoryDecorator(base repositories.UserRepository, tracer trace.Tracer, recordErrors bool) *OpentelemetryTracingUserRepositoryDecorator {
	if base == nil {
		panic("OpentelemetryTracingUserRepositoryDecorator got empty base")
	}
	if tracer == nil {
		panic("OpentelemetryTracingUserRepositoryDecorator got empty tracer")
	}
	return &OpentelemetryTracingUserRepositoryDecorator{
		base:         base,
		tracer:       tracer,
		recordErrors: recordErrors,
	}
}

// Create implements repositories.UserRepository
func (d *OpentelemetryTracingUserRepositoryDecorator) Create(ctx context.Context, user *user.Entity, accountID *accountid.Entity) (userID *userid.Entity, logstash ulo.LogStash, err error) {

	var span trace.Span
	ctx, span = d.tracer.Start(ctx, "/Create")
	defer func() {
		if err != nil && d.recordErrors {
			span.RecordError(err)
		}
		span.End()
	}()
	return d.base.Create(ctx, user, accountID)
}

// DeleteByID implements repositories.UserRepository
func (d *OpentelemetryTracingUserRepositoryDecorator) DeleteByID(ctx context.Context, userID *userid.Entity) (logstash ulo.LogStash, err error) {

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

// GetByID implements repositories.UserRepository
func (d *OpentelemetryTracingUserRepositoryDecorator) GetByID(ctx context.Context, userID *userid.Entity) (user *user.Entity, logstash ulo.LogStash, err error) {

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
