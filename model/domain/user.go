package domain

import "time"

//user with token
type User struct {
	Id         int
	NIM        string
	Email      string
	Name       string
	Password   string
	Role       string
	Created_at time.Time
	Updated_at time.Time
	Token      string
}
