package bot_helpers

import (
	"github.com/BobrePatre/grouse-backend/internal/bot_helpers/delivery"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"Bot Helpers",
	fx.Invoke(delivery.RegisterHelperHandlers),
)
