package web

import "time"

type AnswerResponse struct {
	Id         int                `json:"id"`
	QuestionId int                `json:"question_id"`
	UserId     int                `json:"user_id"`
	Answer     string             `json:"answer"`
	Question   []QuestionResponse `json:"question"`
	User       []UserResponse     `json:"user"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}
