package domain

import "time"

type ResetPassword struct {
	Id         int
	UserId     int
	Token      string
	Expired_at time.Time
	Created_at time.Time
}
