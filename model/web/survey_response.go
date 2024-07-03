package web

import "time"

type SurveyResponse struct {
	Id             int       `json:"id"`
	Title          string    `json:"title"`
	TanggalPosting string    `json:"tanggal_posting"`
	BatasPosting   string    `json:"batas_posting"`
	Created_at     time.Time `json:"created_at"`
	Updated_at     time.Time `json:"updated_at"`
}
