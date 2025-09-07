package telegram

import (
	"calories-tracker/internal/model"
	"context"
	"fmt"
	"log/slog"
)

func (b *TelegramBot) SendMessage(ctx context.Context, u *model.User, m string) error {
	_, err := b.Send(u, m)
	if err != nil {
		return fmt.Errorf("service - SendMessage - b.Bot.Send: %w", err)
	}
	slog.Info("message sent", slog.Int64("user_tid", u.Tid), slog.String("message", m))
	return nil
}
