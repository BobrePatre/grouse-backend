package gateway

import (
	"context"
	"github.com/BobrePatre/grouse-backend/internal/message/dto"
	"github.com/BobrePatre/grouse-backend/internal/message/entity"
	"github.com/BobrePatre/grouse-backend/internal/message/ws"
	"log/slog"
	"strconv"
)

type NotificationTtsGateway struct {
	wsHub  *ws.Hub
	logger *slog.Logger
}

func NewNotificationTtsGateway(wsHub *ws.Hub, logger *slog.Logger) *NotificationTtsGateway {
	return &NotificationTtsGateway{
		wsHub:  wsHub,
		logger: logger,
	}
}

func (g NotificationTtsGateway) NotifyClient(ctx context.Context, notification *entity.Message) error {

	client, err := g.wsHub.GetClient(strconv.FormatInt(notification.UserId, 10))
	if err != nil {
		g.logger.Error("failed to get client", "error", err)
		return err
	}
	sendDto := dto.SendTtsNotificationRequest{
		UserId:  notification.UserId,
		Content: notification.Body,
	}

	err = client.Conn.WriteJSON(sendDto)
	if err != nil {
		g.logger.Error("failed to send tts notification", "error", err)
		return err
	}

	return nil
}
