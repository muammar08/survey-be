package web

import "time"

type AnswerCreateRequest struct {
	SurveyId   int    `json:"survey_id"`
	Answer     string `json:"answer"`
	Created_at time.Time
	Updated_at time.Time
}

type AnswerDeleteRequest struct {
	Id int `json:"id"`
}
