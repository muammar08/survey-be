package service

import (
	"context"
	"database/sql"
	"survey/exception"
	"survey/helper"
	"survey/model/domain"
	"survey/model/web"
	"survey/repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type QuestionServiceImpl struct {
	QuestionRepository repository.QuestionRepository
	SurveyReposityory  repository.SurveyRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewQuestionService(questionRepository repository.QuestionRepository, surveyRepository repository.SurveyRepository, DB *sql.DB, validate *validator.Validate) QuestionService {
	return &QuestionServiceImpl{
		QuestionRepository: questionRepository,
		SurveyReposityory:  surveyRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *QuestionServiceImpl) GetAll(ctx context.Context) []web.QuestionResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	questions := service.QuestionRepository.GetAll(ctx, tx)

	return helper.ToQuestionResponses(questions)
}

func (service *QuestionServiceImpl) AddQuestion(ctx context.Context, requests []web.QuestionCreateRequest) []web.QuestionResponse {
	var questionResponses []web.QuestionResponse

	for _, request := range requests {
		err := service.Validate.Struct(request)
		if err != nil {
			return nil
		}
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return nil
	}
	defer helper.CommitOrRollback(tx)

	requestsBySurveyId := make(map[int][]web.QuestionCreateRequest)
	for _, request := range requests {
		requestsBySurveyId[request.SurveyId] = append(requestsBySurveyId[request.SurveyId], request)
	}

	for surveyID, requestsForSurvey := range requestsBySurveyId {
		survey, err := service.SurveyReposityory.ShowSurvey(ctx, tx, surveyID)
		if err != nil {
			panic(err)
		}

		var questions []domain.Question
		for _, request := range requestsForSurvey {
			question := domain.Question{
				SurveyId:   survey.Id,
				Question:   request.Question,
				Type:       request.Type,
				Created_at: time.Now(),
				Updated_at: time.Now(),
			}
			questions = append(questions, question)
		}

		insertedQuestions, err := service.QuestionRepository.AddQuestion(ctx, tx, questions)
		if err != nil {
			return nil
		}

		for _, question := range insertedQuestions {
			questionResponse := helper.ToQuestionResponse(question)
			questionResponses = append(questionResponses, questionResponse)
		}
	}

	return questionResponses
}

func (service *QuestionServiceImpl) UpdateQuestion(ctx context.Context, request web.QuestionUpdateRequest) web.QuestionResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// question, err := service.QuestionRepository.ShowQuestion(ctx, tx, request.Id)
	// if err != nil {
	// 	panic(exception.NewNotFoundError(err.Error()))
	// }

	questions := []domain.Question{
		{
			Id:         request.Id,
			Question:   request.Question,
			Type:       request.Type,
			Updated_at: time.Now(),
		},
	}

	updatedQuestions, err := service.QuestionRepository.UpdateQuestion(ctx, tx, questions)
	helper.PanicIfError(err)

	return helper.ToQuestionResponse(updatedQuestions[0])
}

func (service *QuestionServiceImpl) DeleteQuestion(ctx context.Context, id int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	question, err := service.QuestionRepository.ShowQuestion(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.QuestionRepository.DeleteQuestion(ctx, tx, question)
}

func (service *QuestionServiceImpl) ShowQuestion(ctx context.Context, id int) web.QuestionResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	question, err := service.QuestionRepository.ShowQuestion(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToQuestionResponse(question)
}

func (service *QuestionServiceImpl) AnswerQuestion(ctx context.Context, id int) web.AnswerQuestion {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	question, err := service.QuestionRepository.AnswerQuestion(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToAnswerQuestionResponse(question)
}
