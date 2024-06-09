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
	questionController controller.QuestionController,
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
	router.POST("/api/survey", surveyController.AddSurvey)
	router.PUT("/api/survey/:surveyId", authMiddleware.ApplyAdminMiddleware(surveyController.UpdateSurvey))
	router.DELETE("/api/survey/:surveyId", authMiddleware.ApplyAdminMiddleware(surveyController.DeleteSurvey))

	//Question
	router.GET("/api/question", questionController.GetAll)
	router.GET("/api/question/:questionId", questionController.ShowQuestion)
	router.POST("/api/question", questionController.AddQuestion)
	router.PUT("/api/question/:questionId", questionController.UpdateQuestion)
	router.DELETE("/api/question/:questionId", questionController.DeleteQuestion)

	//Answer
	router.GET("/api/answer", answerController.GetAll)
	router.GET("/api/answer/:answerId", authMiddleware.ApplyMiddleware(answerController.ShowAnswer))
	router.POST("/api/answer", authMiddleware.ApplyMiddleware(answerController.AddAnswer))
	router.DELETE("/api/answer/:answerId", authMiddleware.ApplyMiddleware(surveyController.DeleteSurvey))

	router.PanicHandler = exception.ErrorHandler

	return router
}
