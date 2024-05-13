package main

import (
	"fmt"
	"net/http"

	"project-workshop/go-api-ecom/app"
	"project-workshop/go-api-ecom/controller"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/repository"
	"project-workshop/go-api-ecom/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	surveyRepository := repository.NewSurveyRepository()
	surveyService := service.NewSurveyService(surveyRepository, db, validate)
	surveyController := controller.NewSurveyController(surveyService)

	answerRepository := repository.NewAnswerRepository()
	answerService := service.NewAnswerService(answerRepository, surveyRepository, userRepository, db, validate)
	answerController := controller.NewAnswerController(answerService)

	router := app.NewRouter(userController, surveyController, answerController)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Adjust the allowed origins according to your needs
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := corsHandler.Handler(router)

	fmt.Println("Server listening on port http://localhost:3000/")

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: handler,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
