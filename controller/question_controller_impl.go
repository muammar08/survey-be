package controller

import (
	"net/http"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/web"
	"project-workshop/go-api-ecom/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type QuestionControllerImpl struct {
	QuestionService service.QuestionService
}

func NewQuestionController(questionService service.QuestionService) QuestionController {
	return &QuestionControllerImpl{
		QuestionService: questionService,
	}
}

func (controller *QuestionControllerImpl) AddQuestion(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var questionCreateRequest []web.QuestionCreateRequest
	helper.ReadFromRequestBody(request, &questionCreateRequest)

	questionResponse := controller.QuestionService.AddQuestion(request.Context(), questionCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   questionResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *QuestionControllerImpl) UpdateQuestion(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	questionUpdateRequest := web.QuestionUpdateRequest{}
	helper.ReadFromRequestBody(request, &questionUpdateRequest)

	questionId, err := strconv.Atoi(params.ByName("questionId"))
	helper.PanicIfError(err)
	questionUpdateRequest.Id = questionId

	questionResponse := controller.QuestionService.UpdateQuestion(request.Context(), questionUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   questionResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *QuestionControllerImpl) DeleteQuestion(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	questionId, err := strconv.Atoi(params.ByName("questionId"))
	helper.PanicIfError(err)

	controller.QuestionService.DeleteQuestion(request.Context(), questionId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *QuestionControllerImpl) ShowQuestion(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	questionId, err := strconv.Atoi(params.ByName("questionId"))
	helper.PanicIfError(err)

	questionResponse := controller.QuestionService.ShowQuestion(request.Context(), questionId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   questionResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *QuestionControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	questionResponse := controller.QuestionService.GetAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   questionResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *QuestionControllerImpl) AnswerQuestion(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	questionId, err := strconv.Atoi(params.ByName("questionId"))
	helper.PanicIfError(err)

	questionResponse := controller.QuestionService.AnswerQuestion(request.Context(), questionId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   questionResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
