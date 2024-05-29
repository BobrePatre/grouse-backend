package message

import (
	"github.com/BobrePatre/grouse-backend/internal/message/delivery"
	"github.com/BobrePatre/grouse-backend/internal/message/gateway"
	"github.com/BobrePatre/grouse-backend/internal/message/interactors"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"Message Service",
	fx.Provide(
		interactors.NewSttInteractor,
		fx.Annotate(
			gateway.NewNotificationGateway,
			fx.As(new(interactors.NotificationGateway)),
		),
	),
	fx.Invoke(
		delivery.RegisterSttHandlers,
	),
)
