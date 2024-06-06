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
	AnswerRepository   repository.AnswerRepository
	QuestionRepository repository.QuestionRepository
	UserRepository     repository.UserRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewAnswerService(answerRepository repository.AnswerRepository, questionRepository repository.QuestionRepository, userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) AnswerService {
	return &AnswerServiceImpl{
		AnswerRepository:   answerRepository,
		QuestionRepository: questionRepository,
		UserRepository:     userRepository,
		DB:                 DB,
		Validate:           validate,
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
	questionIDs := make([]int, len(requests))

	for i, request := range requests {
		err := service.Validate.Struct(request)
		if err != nil {
			return nil
		}
		questionIDs[i] = request.QuestionId
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

	for _, questionID := range questionIDs {
		question, err := service.QuestionRepository.ShowQuestion(ctx, tx, questionID)
		if err != nil {
			panic(err)
		}

		var answers []domain.Answer
		for _, request := range requests {
			if request.QuestionId == questionID {
				answer := domain.Answer{
					QuestionId: question.Id,
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
