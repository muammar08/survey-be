package repository

import (
	"context"
	"database/sql"
	"survey/model/domain"
)

type UserRepository interface {
	Register(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Login(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
	FindByUsername(ctx context.Context, tx *sql.Tx, nim string, email string) (domain.User, error)
	FindByUsernamePublic(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
	FindByRole(ctx context.Context, tx *sql.Tx, role bool) (domain.User, error)
}
