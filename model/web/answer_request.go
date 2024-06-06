package web

import "time"

type AnswerCreateRequest struct {
	QuestionId int    `json:"question_id"`
	Answer     string `json:"answer"`
	Created_at time.Time
	Updated_at time.Time
}

type AnswerDeleteRequest struct {
	Id int `json:"id"`
}
