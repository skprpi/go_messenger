package main

import (
	"context"
	"net"
	"os"
	"os/signal"

	"github.com/YFatMR/go_messenger/auth_service/internal/auth"
	"github.com/YFatMR/go_messenger/auth_service/internal/controllers"
	"github.com/YFatMR/go_messenger/auth_service/internal/repositories"
	"github.com/YFatMR/go_messenger/auth_service/internal/repositories/decorators"
	"github.com/YFatMR/go_messenger/auth_service/internal/repositories/mongorepository"
	"github.com/YFatMR/go_messenger/auth_service/internal/servers"
	"github.com/YFatMR/go_messenger/auth_service/internal/services"
	"github.com/YFatMR/go_messenger/core/pkg/configs/cviper"
	"github.com/YFatMR/go_messenger/core/pkg/loggers"
	"github.com/YFatMR/go_messenger/core/pkg/metrics/prometheus"
	"github.com/YFatMR/go_messenger/core/pkg/mongodb"
	"github.com/YFatMR/go_messenger/core/pkg/traces"
	"github.com/YFatMR/go_messenger/protocol/pkg/proto"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	resourcesdk "go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	config := cviper.New()
	config.AutomaticEnv()

	// Init environment vars
	logLevel := config.GetZapcoreLogLevelRequired("LOG_LEVEL")
	logPath := config.GetStringRequired("LOG_PATH")

	collectDatabaseQueryMetrics := config.GetBoolRequired("ENABLE_DATABASE_QUERY_METRICS")

	jaegerEndpoint := config.GetStringRequired("JAEGER_COLLECTOR_ENDPOINT")
	serviceName := config.GetStringRequired("SERVICE_NAME")
	serviceAddress := config.GetStringRequired("SERVICE_ADDRESS")

	databaseSettings := cviper.NewDatabaseSettingsFromConfig(config)
	metricServiceSettings := cviper.NewMetricMetricServiceSettingsFromConfig(config)

	authTokenSecretKey := config.GetStringRequired("AUTH_TOKEN_SECRET_KEY")
	authTokenExpirationDuration := config.GetSecondsDurationRequired("AUTH_TOKEN_EXPIRATION_SECONDS")

	// Init logger
	zapLogger, err := loggers.NewBaseZapFileLogger(logLevel, logPath)
	if err != nil {
		panic(err)
	}
	logger := loggers.NewOtelZapLoggerWithTraceID(
		otelzap.New(
			zapLogger,
			otelzap.WithTraceIDField(true),
			otelzap.WithMinLevel(zapcore.ErrorLevel),
			otelzap.WithStackTrace(true),
		),
	)
	defer logger.Sync()

	// Init database
	mongoCollection, err := mongodb.Connect(ctx, databaseSettings, logger)
	if err != nil {
		logger.Error("Can't establish connection with database", zap.Error(err))
		panic(err)
	}

	// Init tracing
	logger.Info("Init metrics")
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(jaegerEndpoint)))
	if err != nil {
		panic(err)
	}
	traceProvider, err := traces.NewTraceProvider(exporter, resourcesdk.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(serviceName),
	))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := traceProvider.Shutdown(context.Background()); err != nil { // TODO: Shutdown -> with timeout
			panic(err)
		}
	}()
	tracer := otel.Tracer(serviceName)

	// Init metrics
	if metricServiceSettings != nil {
		go prometheus.ListenAndServeMetrcirService(metricServiceSettings, logger)
	}

	// Init Server
	logger.Info("Init server")
	authManager := auth.NewJWTManager(authTokenSecretKey, authTokenExpirationDuration, logger, tracer)
	passwordValidator := auth.NewDefaultPasswordValidator()

	var repository repositories.AccountRepository
	repository = mongorepository.NewAccountMongoRepository(
		mongoCollection, databaseSettings.GetOperationTimeout(), logger,
	)
	if collectDatabaseQueryMetrics {
		repository = decorators.NewMetricDecorator(repository)
	}

	service := services.NewAccountService(repository, authManager, logger, tracer)
	controller := controllers.NewAccountController(service, passwordValidator, logger, tracer)
	server := servers.NewGRPCAuthServer(controller, logger, tracer)
	s := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
	)

	// Register protobuf server
	logger.Info("Register protobuf server")
	proto.RegisterAuthServer(s, &server)

	// Listen connection
	logger.Info("Server successfully setup. Starting listen...")
	listener, err := net.Listen("tcp", serviceAddress)
	if err != nil {
		panic(err)
	}
	err = s.Serve(listener)
	if err != nil {
		panic(err)
	}
}