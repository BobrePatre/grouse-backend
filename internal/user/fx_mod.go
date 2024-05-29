package user

import (
	"github.com/BobrePatre/grouse-backend/internal/user/delivery"
	"github.com/BobrePatre/grouse-backend/internal/user/repository"
	"github.com/BobrePatre/grouse-backend/internal/user/service/interactors"
	"go.uber.org/fx"
	"log/slog"
)

var Module = fx.Module(
	"User Domain",
	fx.Provide(
		fx.Annotate(
			repository.NewUserRepository,
			fx.As(new(interactors.UserRepository)),
		),
		interactors.NewUserInteractor,
	),

	fx.Invoke(
		delivery.RegisterUserHandlers,
		func(logger *slog.Logger) {
			logger.Info("User Domain Connected")
		},
	),
)
