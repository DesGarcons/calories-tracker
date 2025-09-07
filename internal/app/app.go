package app

import (
	"calories-tracker/config"
	"calories-tracker/internal/adapter/repo"
	"calories-tracker/internal/adapter/telegram"
	"calories-tracker/internal/bothandlers"
	"calories-tracker/internal/common/bot"
	"calories-tracker/internal/service/authorizer"
	"calories-tracker/internal/service/messagesendler"
	"calories-tracker/internal/usecase"
	"calories-tracker/pkg/postgres"
	"fmt"
	"log/slog"
	"time"

	tele "gopkg.in/telebot.v4"
)

type AppEntities struct {
	b *bothandlers.TelegramBot
}

func InitSE(cfg *config.Config) (ae *AppEntities, err error) {
	pref := tele.Settings{
		Token:     cfg.Telegram.Token,
		Poller:    &tele.LongPoller{Timeout: 10 * time.Second},
		ParseMode: "HTML",
	}

	telebot, err := tele.NewBot(pref)
	if err != nil {
		return nil, fmt.Errorf("app - Run - tele.NewBot: %w", err)
	}

	botCommon := bot.New(telebot)

	botSend := telegram.New(botCommon)

	pg, err := postgres.New(cfg.Postgres.Url)
	if err != nil {
		return nil, fmt.Errorf("app - Run - postgres.New: %w", err)
	}

	repo := repo.New(pg)

	messagesendler := messagesendler.New(botSend)
	authorizer := authorizer.New(repo)

	usecase := usecase.New(
		messagesendler,
		authorizer,
	)

	bothandlers := bothandlers.New(botCommon, usecase)

	ae = &AppEntities{
		b: bothandlers,
	}

	return ae, nil
}

func RunApp(ae *AppEntities) {
	logger := slog.Default()
	child := logger.With(
		slog.String("where", "app"),
	)

	slog.SetDefault(child)

	slog.Info("Bot start")

	ae.b.RegisterHandlers()
	ae.b.Start()
}
