package controller

import (
	"net/http"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/web"
	"project-workshop/go-api-ecom/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type SurveyControllerImpl struct {
	SurveyService service.SurveyService
}

func NewSurveyController(surveyService service.SurveyService) SurveyController {
	return &SurveyControllerImpl{
		SurveyService: surveyService,
	}
}

func (controller *SurveyControllerImpl) AddSurvey(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//surveyCreateRequest := web.SurveyCreateRequest{}
	var surveyCreateRequest web.SurveyCreateRequest
	helper.ReadFromRequestBody(request, &surveyCreateRequest)

	surveyResponse := controller.SurveyService.AddSurvey(request.Context(), surveyCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   surveyResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SurveyControllerImpl) UpdateSurvey(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	surveyUpdateRequest := web.SurveyUpdateRequest{}
	helper.ReadFromRequestBody(request, &surveyUpdateRequest)

	surveyId := params.ByName("surveyId")
	id, err := strconv.Atoi(surveyId)
	helper.PanicIfError(err)

	surveyUpdateRequest.Id = id

	surveyResponse := controller.SurveyService.UpdateSurvey(request.Context(), surveyUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   surveyResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SurveyControllerImpl) DeleteSurvey(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	surveyId := params.ByName("surveyId")
	id, err := strconv.Atoi(surveyId)
	helper.PanicIfError(err)

	controller.SurveyService.DeleteSurvey(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SurveyControllerImpl) ShowSurvey(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	surveyId := params.ByName("surveyId")
	id, err := strconv.Atoi(surveyId)
	helper.PanicIfError(err)

	surveyResponse := controller.SurveyService.ShowSurvey(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   surveyResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SurveyControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	surveyResponse := controller.SurveyService.GetAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   surveyResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *SurveyControllerImpl) AllAnswer(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	surveyId := params.ByName("surveyId")
	id, err := strconv.Atoi(surveyId)
	helper.PanicIfError(err)

	surveyResponse := controller.SurveyService.AllAnswer(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   surveyResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
