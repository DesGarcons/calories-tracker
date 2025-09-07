package authorizer

import (
	"calories-tracker/internal/model"
	"fmt"

	"context"
)

type Authorizer struct {
	repo repo
}

type repo interface {
	AddUser(ctx context.Context, u *model.User) error
}

func New(repo repo) *Authorizer {
	return &Authorizer{
		repo: repo,
	}
}

func (a *Authorizer) StartAuth(ctx context.Context, user *model.User) (msg string, err error) {
	err = a.repo.AddUser(ctx, user)
	if err != nil {
		return "Что-то пошло не так", fmt.Errorf("authorizer - StartAuth - a.repo.AddUser: %w", err)
	}
	return "Напишите, пожалуйста ваш вес", nil
}
