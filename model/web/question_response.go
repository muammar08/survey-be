package web

import "time"

type QuestionResponse struct {
	Id         int              `json:"id"`
	SurveyId   int              `json:"survey_id"`
	Question   string           `json:"question"`
	Type       string           `json:"type"`
	Survey     []SurveyResponse `json:"survey"`
	Created_at time.Time        `json:"created_at"`
	Updated_at time.Time        `json:"updated_at"`
}

type AnswerQuestion struct {
	Id       int              `json:"id"`
	SurveyId int              `json:"survey_id"`
	Question string           `json:"question"`
	Type     string           `json:"type"`
	Answer   []AnswerResponse `json:"answer"`
}
