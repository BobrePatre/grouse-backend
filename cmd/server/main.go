package main

import (
	"github.com/BobrePatre/grouse-backend/internal/bot_helpers"
	"github.com/BobrePatre/grouse-backend/internal/infrastructure"
	"github.com/BobrePatre/grouse-backend/internal/message"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"log/slog"
)

func main() {
	app := fx.New(

		infrastructure.Module,
		bot_helpers.Module,
		message.Module,

		fx.WithLogger(func(logger *slog.Logger) fxevent.Logger {
			return &fxevent.SlogLogger{
				Logger: logger,
			}
		}),
	)

	app.Run()
}
