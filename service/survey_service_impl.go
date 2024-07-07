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

type SurveyServiceImpl struct {
	SurveyRepository repository.SurveyRepository
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewSurveyService(surveyRepository repository.SurveyRepository, DB *sql.DB, validate *validator.Validate) SurveyService {
	return &SurveyServiceImpl{
		SurveyRepository: surveyRepository,
		DB:               DB,
		Validate:         validate,
	}
}

func (service *SurveyServiceImpl) AddSurvey(ctx context.Context, request web.SurveyCreateRequest) web.SurveyResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	survey := domain.Survey{
		Title:          request.Title,
		TanggalPosting: request.TanggalPosting,
		BatasPosting:   request.BatasPosting,
		Role:           request.Role,
		Created_at:     time.Now(),
		Updated_at:     time.Now(),
	}

	survey = service.SurveyRepository.AddSurvey(ctx, tx, survey)

	return helper.ToSurveyResponse(survey)
}

func (service *SurveyServiceImpl) UpdateSurvey(ctx context.Context, request web.SurveyUpdateRequest) web.SurveyResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	survey, err := service.SurveyRepository.ShowSurvey(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	survey.Id = request.Id
	survey.Title = request.Title
	survey.TanggalPosting = request.TanggalPosting
	survey.BatasPosting = request.BatasPosting
	survey.Role = request.Role
	survey.Updated_at = time.Now()

	survey = service.SurveyRepository.UpdateSurvey(ctx, tx, survey)

	return helper.ToSurveyResponse(survey)
}

func (service *SurveyServiceImpl) DeleteSurvey(ctx context.Context, id int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	survey, err := service.SurveyRepository.ShowSurvey(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.SurveyRepository.DeleteSurvey(ctx, tx, survey)
}

func (service *SurveyServiceImpl) ShowSurvey(ctx context.Context, id int) web.SurveyResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	survey, err := service.SurveyRepository.ShowSurvey(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToSurveyResponse(survey)
}

func (service *SurveyServiceImpl) GetAll(ctx context.Context) []web.SurveyResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	surveys := service.SurveyRepository.GetAll(ctx, tx)

	return helper.ToSurveyResponses(surveys)
}
