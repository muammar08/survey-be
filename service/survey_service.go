package service

import (
	"context"
	"survey/model/web"
)

type SurveyService interface {
	AddSurvey(ctx context.Context, request web.SurveyCreateRequest) web.SurveyResponse
	UpdateSurvey(ctx context.Context, request web.SurveyUpdateRequest) web.SurveyResponse
	DeleteSurvey(ctx context.Context, id int)
	ShowSurvey(ctx context.Context, id int) web.SurveyResponse
	GetAll(ctx context.Context) []web.SurveyResponse
}
