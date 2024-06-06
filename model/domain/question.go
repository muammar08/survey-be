package domain

import "time"

type Question struct {
	Id         int
	SurveyId   int
	Question   string
	Survey     []Survey
	Created_at time.Time
	Updated_at time.Time
}
