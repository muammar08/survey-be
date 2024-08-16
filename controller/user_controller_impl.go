package controller

import (
	"net/http"
	"survey/helper"
	"survey/model/web"
	"survey/service"

	// "strconv"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse := controller.UserService.Register(request.Context(), userCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userLoginRequest := web.UserLoginRequest{}
	helper.ReadFromRequestBody(request, &userLoginRequest)

	userResponse, err := controller.UserService.Login(request.Context(), userLoginRequest)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) LoginPublic(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userLoginPublicRequest := web.UserLoginPublicRequest{}
	helper.ReadFromRequestBody(request, &userLoginPublicRequest)

	userResponse, err := controller.UserService.LoginPublic(request.Context(), userLoginPublicRequest)
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Error",
			Data:   err.Error(),
		}
		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) SendResetPassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	forgotPasswordRequest := web.ForgotPasswordRequest{}
	helper.ReadFromRequestBody(request, &forgotPasswordRequest)

	userResponse := controller.UserService.SendResetPassword(request.Context(), forgotPasswordRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) VerifyResetPassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	resetPasswordRequest := web.ResetPasswordRequest{}
	helper.ReadFromRequestBody(request, &resetPasswordRequest)

	userResponse := controller.UserService.VerifyResetPassword(request.Context(), resetPasswordRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) ResetPassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	changePasswordRequest := web.ChangePasswordRequest{}
	userId := request.Context().Value("userId").(int)
	helper.ReadFromRequestBody(request, &changePasswordRequest)

	userResponse := controller.UserService.ResetPassword(request.Context(), changePasswordRequest, userId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
