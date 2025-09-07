package repo

import (
	"calories-tracker/internal/model"
	"calories-tracker/pkg/postgres"
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
)

type BotRepo struct {
	*postgres.Postgres
}

// Конструктор
func New(pg *postgres.Postgres) *BotRepo {
	return &BotRepo{pg}
}

func (br *BotRepo) AddUser(ctx context.Context, u *model.User) error {

	tx, err := br.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("BotRepo - AddUser - br.Pool.Begin: %w", err)
	}

	// Добавление пользователя в таблицу users
	sql, args, err := br.Builder.
		Insert("users").Columns("tid").
		Select(
			squirrel.
				Select(fmt.Sprintf("%d", u.Tid)).
				Where("not exists").
				Suffix(fmt.Sprintf("(select id from users where tid = %d)", u.Tid)),
		).ToSql()
	if err != nil {
		return fmt.Errorf("BotRepo - AddUser - br.Builder - users: %w", err)
	}

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("BotRepo - AddUser - tx.Exec: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("BotRepo - AddUser - tx.Commit: %w", err)
	}

	return err
}
