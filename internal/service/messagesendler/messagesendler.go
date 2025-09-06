package messagesendler

import (
	"ccb/internal/model"
	"context"
)

type MessageSender struct {
	messenger messenger
}

type messenger interface {
	SendMessage(ctx context.Context, u *model.User, message string) error
}

func New(messenger messenger) *MessageSender {
	return &MessageSender{
		messenger: messenger,
	}
}

func (ms *MessageSender) SendMessage(ctx context.Context, u *model.User, m string) error {
	return ms.messenger.SendMessage(ctx, u, m)
}
