package service

import (
	"context"
	"database/sql"
	"project-workshop/go-api-ecom/exception"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/domain"
	"project-workshop/go-api-ecom/model/web"
	"project-workshop/go-api-ecom/repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type AnswerServiceImpl struct {
	AnswerRepository repository.AnswerRepository
	SurveyRepository repository.SurveyRepository
	UserRepository   repository.UserRepository
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewAnswerService(answerRepository repository.AnswerRepository, surveyRepository repository.SurveyRepository, userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) AnswerService {
	return &AnswerServiceImpl{
		AnswerRepository: answerRepository,
		SurveyRepository: surveyRepository,
		UserRepository:   userRepository,
		DB:               DB,
		Validate:         validate,
	}
}

func (service *AnswerServiceImpl) GetAll(ctx context.Context) []web.AnswerResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	answers := service.AnswerRepository.GetAll(ctx, tx)

	return helper.ToAnswerResponses(answers)
}

func (service *AnswerServiceImpl) AddAnswer(ctx context.Context, requests []web.AnswerCreateRequest, userId int) []web.AnswerResponse {
	surveyIDs := make([]int, len(requests))

	for i, request := range requests {
		err := service.Validate.Struct(request)
		if err != nil {
			return nil
		}
		surveyIDs[i] = request.SurveyId
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return nil
	}
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(err)
	}

	var answerResponses []web.AnswerResponse

	for _, surveyID := range surveyIDs {
		survey, err := service.SurveyRepository.ShowSurvey(ctx, tx, surveyID)
		if err != nil {
			panic(err)
		}

		var answers []domain.Answer
		for _, request := range requests {
			if request.SurveyId == surveyID {
				answer := domain.Answer{
					SurveyId:   survey.Id,
					UserId:     user.Id,
					Answer:     request.Answer,
					Created_at: time.Now(),
					Updated_at: time.Now(),
				}
				answers = append(answers, answer)
			}
		}

		insertedAnswers, err := service.AnswerRepository.AddAnswer(ctx, tx, answers)
		if err != nil {
			return nil
		}

		for _, answer := range insertedAnswers {
			answerResponse := helper.ToAnswerResponse(answer)
			answerResponses = append(answerResponses, answerResponse)
		}

	}

	return answerResponses
}

func (service *AnswerServiceImpl) DeleteAnswer(ctx context.Context, id int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	answer, err := service.AnswerRepository.ShowAnswer(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.AnswerRepository.DeleteAnswer(ctx, tx, answer)
}

func (service *AnswerServiceImpl) ShowAnswer(ctx context.Context, id int) web.AnswerResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	answer, err := service.AnswerRepository.ShowAnswer(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToAnswerResponse(answer)
}
