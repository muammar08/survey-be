package service

import (
	"context"
	"survey/model/web"
)

type AnswerService interface {
	AddAnswer(ctx context.Context, requests []web.AnswerCreateRequest, userId int) []web.AnswerResponse
	DeleteAnswer(ctx context.Context, id int)
	ShowAnswer(ctx context.Context, id int) web.AnswerResponse
	GetAll(ctx context.Context) []web.AnswerResponse
}
