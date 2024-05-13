package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AnswerController interface {
	AddAnswer(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteAnswer(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	ShowAnswer(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
