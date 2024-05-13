package helper

import (
	"project-workshop/go-api-ecom/model/domain"
	"project-workshop/go-api-ecom/model/web"
	// "strconv"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:       user.Id,
		NIM:      &user.NIM,
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
		Role:     user.Role,
	}
}

func ToSurveyResponse(survey domain.Survey) web.SurveyResponse {
	return web.SurveyResponse{
		Id:         survey.Id,
		Title:      survey.Title,
		Question:   survey.Question,
		Created_at: survey.Created_at,
		Updated_at: survey.Updated_at,
	}
}

func ToAnswerResponse(answer domain.Answer) web.AnswerResponse {
	return web.AnswerResponse{
		Id:         answer.Id,
		SurveyId:   answer.SurveyId,
		UserId:     answer.UserId,
		Answer:     answer.Answer,
		Survey:     ToSurveyResponses(answer.Survey),
		User:       ToUserResponses(answer.User),
		Created_at: answer.Created_at,
		Updated_at: answer.Updated_at,
	}
}

func ToUserResponses(users []domain.User) []web.UserResponse {
	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}

func ToSurveyResponses(surveys []domain.Survey) []web.SurveyResponse {
	var surveyResponses []web.SurveyResponse
	for _, survey := range surveys {
		surveyResponses = append(surveyResponses, ToSurveyResponse(survey))
	}
	return surveyResponses
}

func ToAnswerResponses(answers []domain.Answer) []web.AnswerResponse {
	var answerResponses []web.AnswerResponse
	for _, answer := range answers {
		answerResponses = append(answerResponses, ToAnswerResponse(answer))
	}
	return answerResponses
}
