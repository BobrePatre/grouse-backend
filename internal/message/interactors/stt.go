package interactors

import (
	"context"
	"github.com/BobrePatre/grouse-backend/internal/message/dto"
	"github.com/BobrePatre/grouse-backend/internal/message/entity"
	"log/slog"
)

type NotificationSttGateway interface {
	Notify(ctx context.Context, notification *entity.Message) error
}

type SttInteractor struct {
	notifications NotificationSttGateway
	logger        *slog.Logger
}

func NewSttInteractor(notificationGateway NotificationSttGateway, logger *slog.Logger) *SttInteractor {
	return &SttInteractor{
		notifications: notificationGateway,
		logger:        logger,
	}
}

func (i *SttInteractor) SendNotification(ctx context.Context, dto dto.SendSttNotificationRequest) error {
	notification := entity.NewMessage(
		dto.UserId,
		dto.Content,
	)
	err := i.notifications.Notify(ctx, notification)
	if err != nil {
		slog.Error("failed to send notification", "error", err)
		return err
	}
	slog.Info("notification sent")
	return nil
}
