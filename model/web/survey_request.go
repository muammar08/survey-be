package web

import "time"

type SurveyCreateRequest struct {
	Title      string `json:"title"`
	Created_at time.Time
	Updated_at time.Time
}

type SurveyUpdateRequest struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Updated_at time.Time
}

type SurveyDeleteRequest struct {
	Id int `json:"id"`
}
