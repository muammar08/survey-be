package controller

import (
	"net/http"
	"strconv"
	"survey/helper"
	"survey/model/web"
	"survey/service"

	"github.com/julienschmidt/httprouter"
)

type AnswerControllerImpl struct {
	AnswerService service.AnswerService
}

func NewAnswerController(answerService service.AnswerService) AnswerController {
	return &AnswerControllerImpl{
		AnswerService: answerService,
	}
}

func (controller *AnswerControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	answerResponse := controller.AnswerService.GetAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   answerResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AnswerControllerImpl) AddAnswer(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//surveyCreateRequest := web.SurveyCreateRequest{}
	var answerCreateRequest []web.AnswerCreateRequest
	userId := request.Context().Value("userId").(int)
	helper.ReadFromRequestBody(request, &answerCreateRequest)

	answerResponse := controller.AnswerService.AddAnswer(request.Context(), answerCreateRequest, userId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   answerResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AnswerControllerImpl) DeleteAnswer(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	answerId := params.ByName("answerId")
	id, err := strconv.Atoi(answerId)
	helper.PanicIfError(err)

	controller.AnswerService.DeleteAnswer(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Deleted Success",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AnswerControllerImpl) ShowAnswer(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	answerId := params.ByName("answerId")
	id, err := strconv.Atoi(answerId)
	helper.PanicIfError(err)

	answerResponse := controller.AnswerService.ShowAnswer(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   answerResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
