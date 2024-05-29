package delivery

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
	"log/slog"
)

type BotHelpersHandlers struct {
	logger *slog.Logger
}

func RegisterHelperHandlers(dispatcher *ext.Dispatcher, logger *slog.Logger) error {
	impl := &BotHelpersHandlers{
		logger: logger,
	}

	logger.Info("Registered New handler")
	dispatcher.AddHandler(handlers.NewCommand("start", impl.Start))
	dispatcher.AddHandler(handlers.NewMessage(message.Equal("🆔Узнать ID"), impl.UserId))
	return nil
}

func (h *BotHelpersHandlers) Start(b *gotgbot.Bot, ctx *ext.Context) error {
	kb := gotgbot.ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard: [][]gotgbot.KeyboardButton{
			{
				{
					Text: "🆔Узнать ID",
				},
			},
		},
	}
	_, err := b.SendMessage(ctx.EffectiveChat.Id, "Привет 👋\nВ этот чат я буду присылать то, что тебе говорят.\nЕще я могу помочь тебе узнать твой ID для подключения к приложению ", &gotgbot.SendMessageOpts{
		ReplyMarkup: kb,
	})
	if err != nil {
		h.logger.Error("Failed to send message", "err", err)
	}
	return nil
}

func (h *BotHelpersHandlers) UserId(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := b.SendMessage(ctx.EffectiveChat.Id, fmt.Sprintf("Ваш ID: %d", ctx.EffectiveChat.Id), nil)
	if err != nil {
		return err
	}
	return nil
}
