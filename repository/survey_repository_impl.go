package repository

import (
	"context"
	"database/sql"
	"fmt"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/domain"
	"time"
)

type SurveyRepositoryImpl struct {
}

func NewSurveyRepository() SurveyRepository {
	return &SurveyRepositoryImpl{}
}

func (repository *SurveyRepositoryImpl) AddSurvey(ctx context.Context, tx *sql.Tx, survey domain.Survey) domain.Survey {
	SQL := "INSERT INTO surveys(title, tanggal_posting, batas_posting, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, survey.Title, survey.TanggalPosting, survey.BatasPosting, survey.Created_at, survey.Updated_at)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	survey.Id = int(id)

	return survey
}

func (repository *SurveyRepositoryImpl) UpdateSurvey(ctx context.Context, tx *sql.Tx, survey domain.Survey) domain.Survey {
	SQL := "UPDATE surveys SET title = ?, tanggal_posting = ?, batas_posting = ?, updated_at = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, survey.Title, survey.TanggalPosting, survey.BatasPosting, survey.Updated_at, survey.Id)
	helper.PanicIfError(err)

	return survey
}

func (repository *SurveyRepositoryImpl) DeleteSurvey(ctx context.Context, tx *sql.Tx, survey domain.Survey) domain.Survey {
	SQL := "DELETE FROM surveys where id = ?"
	_, err := tx.ExecContext(ctx, SQL, survey.Id)
	helper.PanicIfError(err)

	return survey
}

func (repository *SurveyRepositoryImpl) ShowSurvey(ctx context.Context, tx *sql.Tx, id int) (domain.Survey, error) {
	SQL := "SELECT * FROM surveys WHERE id = ?"
	rows := tx.QueryRowContext(ctx, SQL, id)
	var survey domain.Survey

	var createdAt, updatedAt []uint8

	err := rows.Scan(&survey.Id, &survey.Title, &survey.TanggalPosting, &survey.BatasPosting, &createdAt, &updatedAt)
	if err != nil {
		return domain.Survey{}, err
	}

	createdAtStr := string(createdAt)
	updatedAtStr := string(updatedAt)

	layout := "2006-01-02 15:04:05"
	createdAtTime, err := time.Parse(layout, createdAtStr)
	if err != nil {
		return domain.Survey{}, err
	}
	updatedAtTime, err := time.Parse(layout, updatedAtStr)
	if err != nil {
		return domain.Survey{}, err
	}

	survey.Created_at = createdAtTime
	survey.Updated_at = updatedAtTime

	return survey, nil
}

func (repository *SurveyRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []domain.Survey {
	SQL := "SELECT * FROM surveys"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var surveys []domain.Survey
	for rows.Next() {
		var survey domain.Survey
		var createdAt, updatedAt string
		err := rows.Scan(&survey.Id, &survey.Title, &survey.TanggalPosting, &survey.BatasPosting, &createdAt, &updatedAt)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return []domain.Survey{}
		}

		createdAtStr := string(createdAt)
		updatedAtStr := string(updatedAt)
		layout := "2006-01-02 15:04:05"
		createdAtTime, err := time.Parse(layout, createdAtStr)
		if err != nil {
			fmt.Println("Error parsing created_at:", err)
			return []domain.Survey{}
		}
		updatedAtTime, err := time.Parse(layout, updatedAtStr)
		if err != nil {
			fmt.Println("Error parsing updated_at:", err)
			return []domain.Survey{}
		}

		survey.Created_at = createdAtTime
		survey.Updated_at = updatedAtTime
		surveys = append(surveys, survey)
	}

	if err = rows.Err(); err != nil {
		//fmt.Println("Error in rows:", err)
		return []domain.Survey{}
	}

	return surveys
}
