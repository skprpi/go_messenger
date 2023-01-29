// Code generated by gowrap. DO NOT EDIT.
// template: ../../../../core/pkg/decorators/templates/repositories/prometheus_metrics.go
// gowrap: http://github.com/hexdigest/gowrap

package decorators

//go:generate gowrap gen -p github.com/YFatMR/go_messenger/user_service/internal/repositories -i UserRepository -t ../../../../core/pkg/decorators/templates/repositories/prometheus_metrics.go -o prometheus_metrics.gen.go -l ""

import (
	"context"
	"time"

	"github.com/YFatMR/go_messenger/core/pkg/errors/cerrors"
	accountid "github.com/YFatMR/go_messenger/user_service/internal/entities/account_id"
	"github.com/YFatMR/go_messenger/user_service/internal/entities/user"
	userid "github.com/YFatMR/go_messenger/user_service/internal/entities/user_id"
	"github.com/YFatMR/go_messenger/user_service/internal/repositories"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Naming rule: https://prometheus.io/docs/practices/naming/

type databaseOperatinTag string

const (
	okStatusTag    = "ok"
	errorStatusTag = "error"
)

var (
	databaseQueryDurationSeconds = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "database_success_query_duration_seconds",
		Buckets: []float64{0.0001, 0.0005, 0.001, 0.005, 0.01, 0.05, 0.1, 0.5, 1, 2, 4},
		Help:    "Duration of one database query",
	}, []string{"status", "operation"})
	databaseQueryProcessedTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "database_query_processed_total",
		Help: "Count of query to database",
	}, []string{"status", "operation"})
	databaseQueryStartProcessTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "database_query_start_process_total",
		Help: "Count of started queries",
	})
)

// PrometheusMetricsUserRepositoryDecorator implements repositories.UserRepository that is instrumented with custom zap logger
type PrometheusMetricsUserRepositoryDecorator struct {
	base repositories.UserRepository
}

// NewPrometheusMetricsUserRepositoryDecorator instruments an implementation of the repositories.UserRepository with simple logging
func NewPrometheusMetricsUserRepositoryDecorator(base repositories.UserRepository) *PrometheusMetricsUserRepositoryDecorator {
	return &PrometheusMetricsUserRepositoryDecorator{
		base: base,
	}
}

// Create implements repositories.UserRepository
func (d *PrometheusMetricsUserRepositoryDecorator) Create(ctx context.Context, user *user.Entity, accountID *accountid.Entity) (userID *userid.Entity, cerr cerrors.Error) {

	startTime := time.Now()
	databaseQueryStartProcessTotal.Inc()
	defer func() {
		functionDuration := time.Since(startTime).Seconds()
		statusTag := errorStatusTag
		if cerr == nil {
			statusTag = okStatusTag
		}
		databaseQueryDurationSeconds.WithLabelValues(statusTag, "create").Observe(functionDuration)
		databaseQueryProcessedTotal.WithLabelValues(statusTag, "create").Inc()
	}()
	return d.base.Create(ctx, user, accountID)

}

// DeleteByID implements repositories.UserRepository
func (d *PrometheusMetricsUserRepositoryDecorator) DeleteByID(ctx context.Context, userID *userid.Entity) (cerr cerrors.Error) {

	startTime := time.Now()
	databaseQueryStartProcessTotal.Inc()
	defer func() {
		functionDuration := time.Since(startTime).Seconds()
		statusTag := errorStatusTag
		if cerr == nil {
			statusTag = okStatusTag
		}
		databaseQueryDurationSeconds.WithLabelValues(statusTag, "deletebyid").Observe(functionDuration)
		databaseQueryProcessedTotal.WithLabelValues(statusTag, "deletebyid").Inc()
	}()
	return d.base.DeleteByID(ctx, userID)

}

// GetByID implements repositories.UserRepository
func (d *PrometheusMetricsUserRepositoryDecorator) GetByID(ctx context.Context, userID *userid.Entity) (user *user.Entity, cerr cerrors.Error) {

	startTime := time.Now()
	databaseQueryStartProcessTotal.Inc()
	defer func() {
		functionDuration := time.Since(startTime).Seconds()
		statusTag := errorStatusTag
		if cerr == nil {
			statusTag = okStatusTag
		}
		databaseQueryDurationSeconds.WithLabelValues(statusTag, "getbyid").Observe(functionDuration)
		databaseQueryProcessedTotal.WithLabelValues(statusTag, "getbyid").Inc()
	}()
	return d.base.GetByID(ctx, userID)

}