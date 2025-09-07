package bothandlers

import (
	"calories-tracker/internal/common/bot"
	"calories-tracker/internal/model"
	"context"
	"fmt"
	"log/slog"

	tele "gopkg.in/telebot.v4"
)

type TelegramBot struct {
	*bot.Bot
	uc usecase
}

type usecase interface {
	StartAuth(ctx context.Context, user *model.User) (msg string, err error)
}

func New(b *bot.Bot, uc usecase) *TelegramBot {

	return &TelegramBot{
		b,
		uc,
	}
}

func (tb *TelegramBot) RegisterHandlers() {
	tb.Bot.Handle("/start", func(c tele.Context) error {
		return c.Send("Привет! Я помогу тебе следить за калориями\n\nДля начала нажми кнопку ниже", tb.Bot.Menus.AuthMenu)
	})

	tb.Bot.Handle(tb.Bot.Buttons.AuthorizationBtn, func(c tele.Context) error {
		ctx := context.Background()
		user := model.User{
			Tid: c.Sender().ID,
		}
		msg, err := tb.uc.StartAuth(ctx, &user)
		if err != nil {
			slog.Error(fmt.Sprintln(err))

		}
		return c.Send(msg)
	})
}
