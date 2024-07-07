package helper

import (
	"survey/model/domain"
	"survey/model/web"
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
		Id:             survey.Id,
		Title:          survey.Title,
		TanggalPosting: survey.TanggalPosting,
		BatasPosting:   survey.BatasPosting,
		Role:           survey.Role,
		Created_at:     survey.Created_at,
		Updated_at:     survey.Updated_at,
	}
}

func ToQuestionResponse(question domain.Question) web.QuestionResponse {
	return web.QuestionResponse{
		Id:         question.Id,
		SurveyId:   question.SurveyId,
		Question:   question.Question,
		Type:       question.Type,
		Survey:     ToSurveyResponses(question.Survey),
		Created_at: question.Created_at,
		Updated_at: question.Updated_at,
	}
}

func ToAnswerQuestionResponse(question domain.AnswerQuestion) web.AnswerQuestion {
	return web.AnswerQuestion{
		Id:       question.Id,
		SurveyId: question.SurveyId,
		Question: question.Question,
		Type:     question.Type,
		Answer:   ToAnswerResponses(question.Answer),
	}
}

func ToAnswerResponse(answer domain.Answer) web.AnswerResponse {
	return web.AnswerResponse{
		Id:         answer.Id,
		QuestionId: answer.QuestionId,
		UserId:     answer.UserId,
		Answer:     answer.Answer,
		Question:   ToQuestionResponses(answer.Question),
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

func ToQuestionResponses(questions []domain.Question) []web.QuestionResponse {
	var questionResponses []web.QuestionResponse
	for _, question := range questions {
		questionResponses = append(questionResponses, ToQuestionResponse(question))
	}
	return questionResponses
}

func ToAnswerResponses(answers []domain.Answer) []web.AnswerResponse {
	var answerResponses []web.AnswerResponse
	for _, answer := range answers {
		answerResponses = append(answerResponses, ToAnswerResponse(answer))
	}
	return answerResponses
}
