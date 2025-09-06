package telegram

import "ccb/internal/common/bot"

type TelegramBot struct {
	*bot.Bot
}

func New(b *bot.Bot) *TelegramBot {
	return &TelegramBot{
		b,
	}
}
