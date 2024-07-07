package web

import "time"

type SurveyCreateRequest struct {
	Title          string `json:"title"`
	TanggalPosting string `json:"tanggal_posting"`
	BatasPosting   string `json:"batas_posting"`
	Role           string `json:"role"`
	Created_at     time.Time
	Updated_at     time.Time
}

type SurveyUpdateRequest struct {
	Id             int    `json:"id"`
	Title          string `json:"title"`
	TanggalPosting string `json:"tanggal_posting"`
	BatasPosting   string `json:"batas_posting"`
	Role           string `json:"role"`
	Updated_at     time.Time
}

type SurveyDeleteRequest struct {
	Id int `json:"id"`
}
