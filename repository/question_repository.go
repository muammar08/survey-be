package repository

import (
	"context"
	"database/sql"
	"project-workshop/go-api-ecom/model/domain"
)

type QuestionRepository interface {
	AddQuestion(ctx context.Context, tx *sql.Tx, questions []domain.Question) ([]domain.Question, error)
	UpdateQuestion(ctx context.Context, tx *sql.Tx, question []domain.Question) ([]domain.Question, error)
	DeleteQuestion(ctx context.Context, tx *sql.Tx, question domain.Question) domain.Question
	ShowQuestion(ctx context.Context, tx *sql.Tx, id int) (domain.Question, error)
	GetAll(ctx context.Context, tx *sql.Tx) []domain.Question
}
