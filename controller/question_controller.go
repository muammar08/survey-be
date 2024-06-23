package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type QuestionController interface {
	AddQuestion(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateQuestion(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteQuestion(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	ShowQuestion(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	AnswerQuestion(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
