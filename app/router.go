package app

import (
	// "net/http"
	"project-workshop/go-api-ecom/controller"
	"project-workshop/go-api-ecom/exception"
	"project-workshop/go-api-ecom/middleware"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(
	userController controller.UserController,
	surveyController controller.SurveyController,
	answerController controller.AnswerController) *httprouter.Router {
	router := httprouter.New()

	// Middleware
	authMiddleware := middleware.Middleware{}

	// Auth
	router.POST("/api/register", userController.Register)
	router.POST("/api/login", userController.Login)
	router.POST("/api/loginpublic", userController.LoginPublic)

	//Survey
	router.GET("/api/survey", authMiddleware.ApplyAdminMiddleware(surveyController.GetAll))
	router.GET("/api/survey/:surveyId", authMiddleware.ApplyAdminMiddleware(surveyController.ShowSurvey))
	router.POST("/api/addsurvey", authMiddleware.ApplyAdminMiddleware(surveyController.AddSurvey))
	router.PUT("/api/updatesurvey/:surveyId", authMiddleware.ApplyAdminMiddleware(surveyController.UpdateSurvey))
	router.DELETE("/api/deletesurvey/:surveyId", authMiddleware.ApplyAdminMiddleware(surveyController.DeleteSurvey))

	//Answer
	router.GET("/api/answer", authMiddleware.ApplyMiddleware(answerController.GetAll))
	router.GET("/api/answer/:answerId", authMiddleware.ApplyMiddleware(answerController.ShowAnswer))
	router.POST("/api/addanswer", authMiddleware.ApplyMiddleware(answerController.AddAnswer))
	router.DELETE("/api/deleteanswer/:answerId", authMiddleware.ApplyMiddleware(surveyController.DeleteSurvey))

	router.PanicHandler = exception.ErrorHandler

	return router
}
