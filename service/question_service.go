package service

import (
	"context"
	"survey/model/web"
)

type QuestionService interface {
	AddQuestion(ctx context.Context, requests []web.QuestionCreateRequest) []web.QuestionResponse
	UpdateQuestion(ctx context.Context, request web.QuestionUpdateRequest) web.QuestionResponse
	DeleteQuestion(ctx context.Context, id int)
	ShowQuestion(ctx context.Context, id int) web.QuestionResponse
	AnswerQuestion(ctx context.Context, id int) web.AnswerQuestion
	GetAll(ctx context.Context) []web.QuestionResponse
}
