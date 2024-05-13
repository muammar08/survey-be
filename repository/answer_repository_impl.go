package repository

import (
	"context"
	"database/sql"
	"fmt"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/domain"
	"time"
)

type AnswerRepositoryImpl struct {
}

func NewAnswerRepository() AnswerRepository {
	return &AnswerRepositoryImpl{}
}

func (repository *AnswerRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []domain.Answer {
	SQL := `
		SELECT a.id, a.survey_id, a.user_id, a.answer, a.created_at, a.updated_at, s.id, s.title, s.question, u.id, u.nim, u.email, u.name
		FROM answers a
		JOIN surveys s ON a.survey_id = s.id
		JOIN users u ON a.user_id = u.id
	`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var answers []domain.Answer
	for rows.Next() {
		var answer domain.Answer
		var survey domain.Survey
		var user domain.User
		var createdAt, updatedAt []uint8
		var nimPtr *string

		err := rows.Scan(&answer.Id, &answer.SurveyId, &answer.UserId, &answer.Answer, &createdAt, &updatedAt, &survey.Id, &survey.Title, &survey.Question, &user.Id, &nimPtr, &user.Email, &user.Name)
		if err != nil {
			return []domain.Answer{}
		}

		createdAtStr := string(createdAt)
		updatedAtStr := string(updatedAt)
		layout := "2006-01-02 15:04:05"
		createdAtTime, err := time.Parse(layout, createdAtStr)
		if err != nil {
			fmt.Println("Error parsing created_at:", err)
			return []domain.Answer{}
		}
		updatedAtTime, err := time.Parse(layout, updatedAtStr)
		if err != nil {
			fmt.Println("Error parsing updated_at:", err)
			return []domain.Answer{}
		}

		answer.Created_at = createdAtTime
		answer.Updated_at = updatedAtTime

		if nimPtr != nil {
			user.NIM = *nimPtr
		}

		answer.Survey = append(answer.Survey, survey)
		answer.User = append(answer.User, user)
		answers = append(answers, answer)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Error in rows:", err)
		return []domain.Answer{}
	}

	return answers
}

func (repository *AnswerRepositoryImpl) AddAnswer(ctx context.Context, tx *sql.Tx, answers []domain.Answer) ([]domain.Answer, error) {
	SQL := "INSERT INTO answers(survey_id, user_id, answer, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"

	var insertedAnswer []domain.Answer

	for _, answer := range answers {
		result, err := tx.ExecContext(ctx, SQL, answer.SurveyId, answer.UserId, answer.Answer, answer.Created_at, answer.Updated_at)
		if err != nil {
			return nil, err
		}

		id, err := result.LastInsertId()
		if err != nil {
			return nil, err
		}

		answer.Id = int(id)
		insertedAnswer = append(insertedAnswer, answer)
	}

	return insertedAnswer, nil
}

func (repository *AnswerRepositoryImpl) DeleteAnswer(ctx context.Context, tx *sql.Tx, answer domain.Answer) domain.Answer {
	SQL := "DELETE FROM answers WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, answer.Id)
	helper.PanicIfError(err)

	return answer
}

func (repository *AnswerRepositoryImpl) ShowAnswer(ctx context.Context, tx *sql.Tx, id int) (domain.Answer, error) {
	SQL := `
        SELECT a.id, a.survey_id, a.user_id, a.answer, a.created_at, a.updated_at,
               s.id AS survey_id, s.title AS survey_title, s.question AS survey_question,
               u.id AS user_id, u.nim AS user_nim, u.email AS user_email, u.name AS user_name
        FROM answers a
        JOIN surveys s ON a.survey_id = s.id
        JOIN users u ON a.user_id = u.id
        WHERE a.id = ?
    `
	rows := tx.QueryRowContext(ctx, SQL, id)
	var answer domain.Answer
	var survey domain.Survey
	var user domain.User

	var createdAt, updatedAt []uint8
	var userNIM *string

	err := rows.Scan(&answer.Id, &answer.SurveyId, &answer.UserId, &answer.Answer, &createdAt, &updatedAt,
		&survey.Id, &survey.Title, &survey.Question, &user.Id, &userNIM, &user.Email, &user.Name)
	if err != nil {
		return domain.Answer{}, err
	}

	createdAtStr := string(createdAt)
	updatedAtStr := string(updatedAt)

	layout := "2006-01-02 15:04:05"
	createdAtTime, err := time.Parse(layout, createdAtStr)
	if err != nil {
		return domain.Answer{}, err
	}
	updatedAtTime, err := time.Parse(layout, updatedAtStr)
	if err != nil {
		return domain.Answer{}, err
	}

	answer.Created_at = createdAtTime
	answer.Updated_at = updatedAtTime

	if userNIM != nil {
		user.NIM = *userNIM
	}

	answer.Survey = append(answer.Survey, survey)
	answer.User = append(answer.User, user)

	return answer, nil
}
