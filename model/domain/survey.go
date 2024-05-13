package domain

import "time"

type Survey struct {
	Id         int
	Title      string
	Question   string
	Created_at time.Time
	Updated_at time.Time
}
