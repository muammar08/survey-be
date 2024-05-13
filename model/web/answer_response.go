package web

import "time"

type AnswerResponse struct {
	Id         int              `json:"id"`
	SurveyId   int              `json:"survey_id"`
	UserId     int              `json:"user_id"`
	Answer     string           `json:"answer"`
	Survey     []SurveyResponse `json:"survey"`
	User       []UserResponse   `json:"user"`
	Created_at time.Time        `json:"created_at"`
	Updated_at time.Time        `json:"updated_at"`
}
