package repository

import (
	"context"
	"database/sql"
	"survey/model/domain"
)

type SurveyRepository interface {
	AddSurvey(ctx context.Context, tx *sql.Tx, survey domain.Survey) domain.Survey
	UpdateSurvey(ctx context.Context, tx *sql.Tx, survey domain.Survey) domain.Survey
	DeleteSurvey(ctx context.Context, tx *sql.Tx, survey domain.Survey) domain.Survey
	ShowSurvey(ctx context.Context, tx *sql.Tx, id int) (domain.Survey, error)
	GetAll(ctx context.Context, tx *sql.Tx) []domain.Survey
}
