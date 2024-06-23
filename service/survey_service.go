package service

import (
	"context"
	"project-workshop/go-api-ecom/model/web"
)

type SurveyService interface {
	AddSurvey(ctx context.Context, request web.SurveyCreateRequest) web.SurveyResponse
	UpdateSurvey(ctx context.Context, request web.SurveyUpdateRequest) web.SurveyResponse
	DeleteSurvey(ctx context.Context, id int)
	ShowSurvey(ctx context.Context, id int) web.SurveyResponse
	AllAnswer(ctx context.Context, id int) web.AllAnswerResponse
	GetAll(ctx context.Context) []web.SurveyResponse
}
