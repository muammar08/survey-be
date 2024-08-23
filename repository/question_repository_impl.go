package repository

import (
	"context"
	"database/sql"
	"fmt"
	"survey/helper"
	"survey/model/domain"
	"time"
)

type QuestionRepositoryImpl struct {
}

func NewQuestionRepository() QuestionRepository {
	return &QuestionRepositoryImpl{}
}

func (repository *QuestionRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []domain.Question {
	SQL := `
		SELECT q.id, q.survey_id, q.question, q.type, q.created_at, q.updated_at, s.id, s.title
		FROM questions q
		JOIN surveys s ON q.survey_id = s.id
	`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var questions []domain.Question
	for rows.Next() {
		var question domain.Question
		var survey domain.Survey
		var createdAt, updatedAt []uint8

		err := rows.Scan(&question.Id, &question.SurveyId, &question.Question, &question.Type, &createdAt, &updatedAt, &survey.Id, &survey.Title)
		if err != nil {
			return []domain.Question{}
		}

		createdAtStr := string(createdAt)
		updatedAtStr := string(updatedAt)
		layout := "2006-01-02 15:04:05"
		createdAtTime, err := time.Parse(layout, createdAtStr)
		if err != nil {
			fmt.Println("Error parsing created_at:", err)
			return []domain.Question{}
		}
		updatedAtTime, err := time.Parse(layout, updatedAtStr)
		if err != nil {
			fmt.Println("Error parsing updated_at:", err)
			return []domain.Question{}
		}

		question.Created_at = createdAtTime
		question.Updated_at = updatedAtTime

		question.Survey = append(question.Survey, survey)
		questions = append(questions, question)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Error in rows:", err)
		return []domain.Question{}
	}

	return questions
}

func (repository *QuestionRepositoryImpl) AddQuestion(ctx context.Context, tx *sql.Tx, questions []domain.Question) ([]domain.Question, error) {
	SQL := "INSERT INTO questions(survey_id, question, type, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"

	var insertedQuestion []domain.Question

	for _, question := range questions {
		result, err := tx.ExecContext(ctx, SQL, question.SurveyId, question.Question, question.Type, question.Created_at, question.Updated_at)
		if err != nil {
			return nil, err
		}

		id, err := result.LastInsertId()
		if err != nil {
			return nil, err
		}

		question.Id = int(id)
		insertedQuestion = append(insertedQuestion, question)
	}

	return insertedQuestion, nil
}

func (repository *QuestionRepositoryImpl) UpdateQuestion(ctx context.Context, tx *sql.Tx, questions []domain.Question) ([]domain.Question, error) {
	SQL := "UPDATE questions SET question = ?, type = ?, updated_at = ? where id = ?"

	var updateQuestion []domain.Question

	for _, question := range questions {
		_, err := tx.ExecContext(ctx, SQL, question.Question, question.Type, question.Updated_at, question.Id)
		if err != nil {
			fmt.Println("err", err)
		}

		updateQuestion = append(updateQuestion, question)
	}

	return updateQuestion, nil
}

func (repository *QuestionRepositoryImpl) DeleteQuestion(ctx context.Context, tx *sql.Tx, question domain.Question) domain.Question {
	SQL := "DELETE FROM questions WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, question.Id)
	helper.PanicIfError(err)

	return question
}

func (repository *QuestionRepositoryImpl) ShowQuestion(ctx context.Context, tx *sql.Tx, id int) (domain.Question, error) {
	SQL := `
		SELECT q.id, q.survey_id, q.question, q.type, q.created_at, q.updated_at, s.id, s.title
		FROM questions q
		JOIN surveys s ON q.survey_id = s.id
		WHERE q.id = ?
	`
	rows := tx.QueryRowContext(ctx, SQL, id)
	var question domain.Question
	var survey domain.Survey
	var createdAt, updatedAt []uint8

	err := rows.Scan(&question.Id, &question.SurveyId, &question.Question, &question.Type, &createdAt, &updatedAt, &survey.Id, &survey.Title)
	if err != nil {
		return domain.Question{}, err
	}

	createdAtStr := string(createdAt)
	updatedAtStr := string(updatedAt)

	layout := "2006-01-02 15:04:05"
	createdAtTime, err := time.Parse(layout, createdAtStr)
	if err != nil {
		return domain.Question{}, err
	}
	updatedAtTime, err := time.Parse(layout, updatedAtStr)
	if err != nil {
		return domain.Question{}, err
	}

	question.Created_at = createdAtTime
	question.Updated_at = updatedAtTime

	question.Survey = append(question.Survey, survey)

	return question, nil
}

func (repository *QuestionRepositoryImpl) AnswerQuestion(ctx context.Context, tx *sql.Tx, id int) (domain.AnswerQuestion, error) {
	SQL := `
		SELECT q.id, q.survey_id, q.question, q.type, 
		       a.id, a.question_id, a.user_id, a.answer, 
		       u.id, u.email, u.name
		FROM questions q
		JOIN answers a ON q.id = a.question_id
		JOIN users u ON a.user_id = u.id
		WHERE q.id = ?
	`

	rows, err := tx.QueryContext(ctx, SQL, id)
	if err != nil {
		return domain.AnswerQuestion{}, err
	}
	defer rows.Close()

	var question domain.AnswerQuestion
	question.Answer = make([]domain.Answer, 0)

	answerMap := make(map[int]domain.Answer)

	for rows.Next() {
		var answer domain.Answer
		var user domain.User

		err := rows.Scan(&question.Id, &question.SurveyId, &question.Question, &question.Type,
			&answer.Id, &answer.QuestionId, &answer.UserId, &answer.Answer,
			&user.Id, &user.Email, &user.Name)

		if err != nil {
			fmt.Println("err", err)
		}

		answer.User = append(answer.User, user)
		answerMap[answer.Id] = answer
	}

	for _, answer := range answerMap {
		question.Answer = append(question.Answer, answer)
	}

	return question, nil
}
