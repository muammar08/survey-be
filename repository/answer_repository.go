package repository

import (
	"context"
	"database/sql"
	"survey/model/domain"
)

type AnswerRepository interface {
	AddAnswer(ctx context.Context, tx *sql.Tx, answers []domain.Answer) ([]domain.Answer, error)
	DeleteAnswer(ctx context.Context, tx *sql.Tx, answer domain.Answer) domain.Answer
	ShowAnswer(ctx context.Context, tx *sql.Tx, id int) (domain.Answer, error)
	GetAll(ctx context.Context, tx *sql.Tx) []domain.Answer
}
