package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type SurveyController interface {
	AddSurvey(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateSurvey(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteSurvey(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	ShowSurvey(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	AllAnswer(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
