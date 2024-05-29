package gateway

import (
	"context"
	"github.com/BobrePatre/grouse-backend/internal/message/entity"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"log/slog"
)

type NotificationGateway struct {
	tgBot  *gotgbot.Bot
	logger *slog.Logger
}

func NewNotificationGateway(tgBot *gotgbot.Bot, logger *slog.Logger) *NotificationGateway {
	return &NotificationGateway{
		tgBot:  tgBot,
		logger: logger,
	}
}

func (g NotificationGateway) Notify(ctx context.Context, notification *entity.Message) error {

	_, err := g.tgBot.SendMessage(
		notification.ReciverId,
		notification.Body,
		nil,
	)
	if err != nil {
		g.logger.Error("failed to send notification", "error", err)
		return err
	}

	return nil
}
