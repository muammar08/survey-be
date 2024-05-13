package web

import "time"

type SurveyCreateRequest struct {
	Title      string `json:"title"`
	Question   string `json:"question"`
	Created_at time.Time
	Updated_at time.Time
}

type SurveyUpdateRequest struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Question   string `json:"question"`
	Updated_at time.Time
}

type SurveyDeleteRequest struct {
	Id int `json:"id"`
}
