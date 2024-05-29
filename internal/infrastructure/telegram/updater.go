package telegram

import (
	"context"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"go.uber.org/fx"
	"log/slog"
)

func NewUpdater(dispatcher *ext.Dispatcher) *ext.Updater {
	return ext.NewUpdater(dispatcher, nil)
}

func StartPolling(lc fx.Lifecycle, bot *gotgbot.Bot, updater *ext.Updater, logger *slog.Logger) error {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("starting telegram bot")
			err := updater.StartPolling(bot, &ext.PollingOpts{
				DropPendingUpdates: true,
			})
			if err != nil {
				return err
			}
			go func() {
				updater.Idle()
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return updater.Stop()
		},
	})
	return nil
}
