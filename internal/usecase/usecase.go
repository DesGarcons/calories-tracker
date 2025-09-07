package usecase

import (
	"calories-tracker/internal/model"
	"context"
)

type Usecase struct {
	messageSender messageSender
	authorizer    authorizer
}

type messageSender interface {
	SendMessage(ctx context.Context, u *model.User, m string) error
}

type authorizer interface {
	StartAuth(ctx context.Context, user *model.User) (msg string, err error)
}

func New(messageSender messageSender, authorizer authorizer) *Usecase {
	return &Usecase{
		messageSender: messageSender,
		authorizer:    authorizer,
	}
}

func (uc *Usecase) StartAuth(ctx context.Context, user *model.User) (msg string, err error) {
	return uc.authorizer.StartAuth(ctx, user)
}
