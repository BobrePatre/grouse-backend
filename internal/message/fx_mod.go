package message

import (
	"github.com/BobrePatre/grouse-backend/internal/message/delivery"
	"github.com/BobrePatre/grouse-backend/internal/message/gateway"
	"github.com/BobrePatre/grouse-backend/internal/message/interactors"
	"github.com/BobrePatre/grouse-backend/internal/message/ws"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"Message Service",
	fx.Provide(
		interactors.NewSttInteractor,
		interactors.NewTtsInteractor,
		ws.NewHub,
		fx.Annotate(
			gateway.NewNotificationSttGateway,
			fx.As(new(interactors.NotificationSttGateway)),
		),
		fx.Annotate(
			gateway.NewNotificationTtsGateway,
			fx.As(new(interactors.NotificationTtsGateway)),
		),
	),
	fx.Invoke(
		ws.RegisterWsHub,
		delivery.RegisterTtsHandlers,
		delivery.RegisterSttHandlers,
	),
)
