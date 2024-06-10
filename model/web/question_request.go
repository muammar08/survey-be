package web

import "time"

type QuestionCreateRequest struct {
	SurveyId   int    `json:"survey_id"`
	Question   string `json:"question"`
	Type       string `json:"type"`
	Created_at time.Time
	Updated_at time.Time
}

type QuestionUpdateRequest struct {
	Id         int    `json:"id"`
	Question   string `json:"question"`
	Type       string `json:"type"`
	Updated_at time.Time
}

type QuestionDeleteRequest struct {
	Id int `json:"id"`
}
