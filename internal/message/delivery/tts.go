package delivery

import (
	"context"
	"github.com/BobrePatre/grouse-backend/internal/message/dto"
	"github.com/BobrePatre/grouse-backend/internal/message/interactors"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
	"log/slog"
)

type TtsHandlers struct {
	logger     *slog.Logger
	interactor *interactors.TtsInteractor
}

func RegisterTtsHandlers(dispatcher *ext.Dispatcher, ttsInteractor *interactors.TtsInteractor, logger *slog.Logger) error {
	impl := &TtsHandlers{
		logger:     logger,
		interactor: ttsInteractor,
	}

	logger.Info("Registered New handler")
	dispatcher.AddHandler(handlers.NewMessage(message.Text, impl.OnMessage))
	return nil
}

func (h *TtsHandlers) OnMessage(b *gotgbot.Bot, ctx *ext.Context) error {

	c := context.Background()
	err := h.interactor.SendNotification(c, dto.SendTtsNotificationRequest{
		UserId:  ctx.EffectiveChat.Id,
		Content: ctx.Message.Text,
	})
	if err != nil {
		_, err := b.SendMessage(ctx.EffectiveChat.Id, "Системе не удалось отправить сообщение в мобильное приложение 😭", nil)
		if err != nil {
			h.logger.Error("Failed to send message", "error", err)
			return err
		}
		return err
	}

	return nil
}
