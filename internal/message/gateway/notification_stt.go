package gateway

import (
	"context"
	"github.com/BobrePatre/grouse-backend/internal/message/entity"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"log/slog"
)

type NotificationSttGateway struct {
	tgBot  *gotgbot.Bot
	logger *slog.Logger
}

func NewNotificationSttGateway(tgBot *gotgbot.Bot, logger *slog.Logger) *NotificationSttGateway {
	return &NotificationSttGateway{
		tgBot:  tgBot,
		logger: logger,
	}
}

func (g NotificationSttGateway) Notify(ctx context.Context, notification *entity.Message) error {

	_, err := g.tgBot.SendMessage(
		notification.UserId,
		notification.Body,
		nil,
	)
	if err != nil {
		g.logger.Error("failed to send notification", "error", err)
		return err
	}

	return nil
}
