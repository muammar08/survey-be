package domain

import "time"

type Answer struct {
	Id         int
	SurveyId   int
	UserId     int
	Answer     string
	Survey     []Survey
	User       []User
	Created_at time.Time
	Updated_at time.Time
}
