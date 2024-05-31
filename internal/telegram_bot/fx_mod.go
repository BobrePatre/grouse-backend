package telegram_bot

import (
	"github.com/BobrePatre/grouse-backend/internal/telegram_bot/handlers"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"Bot Helpers",
	fx.Invoke(handlers.RegisterHelperHandlers),
)
