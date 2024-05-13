package repository

import (
	"context"
	"database/sql"
	"project-workshop/go-api-ecom/model/domain"
)

type SurveyRepository interface {
	AddSurvey(ctx context.Context, tx *sql.Tx, survey []domain.Survey) ([]domain.Survey, error)
	UpdateSurvey(ctx context.Context, tx *sql.Tx, survey domain.Survey) domain.Survey
	DeleteSurvey(ctx context.Context, tx *sql.Tx, survey domain.Survey) domain.Survey
	ShowSurvey(ctx context.Context, tx *sql.Tx, id int) (domain.Survey, error)
	GetAll(ctx context.Context, tx *sql.Tx) []domain.Survey
}
