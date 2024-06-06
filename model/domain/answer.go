package domain

import "time"

type Answer struct {
	Id         int
	QuestionId int
	UserId     int
	Answer     string
	Question   []Question
	User       []User
	Created_at time.Time
	Updated_at time.Time
}
