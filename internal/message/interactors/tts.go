package interactors

import (
	"context"
	"github.com/BobrePatre/grouse-backend/internal/message/dto"
	"github.com/BobrePatre/grouse-backend/internal/message/entity"
	"log/slog"
)

type NotificationTtsGateway interface {
	NotifyClient(ctx context.Context, notification *entity.Message) error
}

type TtsInteractor struct {
	logger        *slog.Logger
	notifyGateway NotificationTtsGateway
}

func NewTtsInteractor(logger *slog.Logger, gateway NotificationTtsGateway) *TtsInteractor {
	return &TtsInteractor{
		logger:        logger,
		notifyGateway: gateway,
	}
}

func (i *TtsInteractor) SendNotification(ctx context.Context, dto dto.SendTtsNotificationRequest) error {
	message := entity.NewMessage(
		dto.UserId,
		dto.Content,
	)
	err := i.notifyGateway.NotifyClient(ctx, message)
	if err != nil {
		i.logger.Error("failed to send tts notification", "error", err)
		return err
	}
	return nil
}
