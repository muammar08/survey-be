package domain

import "time"

type Survey struct {
	Id             int
	Title          string
	TanggalPosting string
	BatasPosting   string
	Created_at     time.Time
	Updated_at     time.Time
}
