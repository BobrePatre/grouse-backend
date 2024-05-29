package infrastructure

import (
	"github.com/BobrePatre/grouse-backend/internal/infrastructure/grpc"
	"github.com/BobrePatre/grouse-backend/internal/infrastructure/http"
	"github.com/BobrePatre/grouse-backend/internal/infrastructure/logging"
	"github.com/BobrePatre/grouse-backend/internal/infrastructure/validate"
	"go.uber.org/fx"
	"log/slog"
)

var Module = fx.Module(
	"Infrastructure",

	// Validator
	fx.Provide(
		validate.NewValidate,
	),

	// Logger
	fx.Provide(
		logging.LoadConfig,
		logging.Logger,
	),

	// Grpc
	fx.Provide(
		grpc.LoadConfig,
		grpc.AsUnaryServerInterceptor(grpc.NewLoggingInterceptor),
		fx.Annotate(
			grpc.NewGrpcServer,
			fx.ParamTags("", grpc.UnaryServerInterceptorGroup),
		),
	),

	// Http
	fx.Provide(
		http.LoadConfig,
		http.NewGatewayServer,
		http.NewHttpServer,
	),

	// Module Entrypoint
	fx.Invoke(
		grpc.RunGrpcServer,
		http.RunHttpServer,
		func(logger *slog.Logger) {
			logger.Info("Infrastructure Initialized")
		},
	),
)
