package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"

	"survey/app"
	"survey/controller"
	"survey/helper"
	"survey/repository"
	"survey/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db := app.NewDB()
	validate := validator.New()
	smtpAuth := smtp.PlainAuth("PakThani", os.Getenv("SMTP_EMAIL"), os.Getenv("SMTP_PASSWORD"), os.Getenv("SMTP_HOST"))
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate, smtpAuth, smtpHost, smtpPort)
	userController := controller.NewUserController(userService)

	surveyRepository := repository.NewSurveyRepository()
	surveyService := service.NewSurveyService(surveyRepository, db, validate)
	surveyController := controller.NewSurveyController(surveyService)

	questionRepository := repository.NewQuestionRepository()
	questionService := service.NewQuestionService(questionRepository, surveyRepository, db, validate)
	questionController := controller.NewQuestionController(questionService)

	answerRepository := repository.NewAnswerRepository()
	answerService := service.NewAnswerService(answerRepository, questionRepository, userRepository, db, validate)
	answerController := controller.NewAnswerController(answerService)

	router := app.NewRouter(userController, surveyController, questionController, answerController)

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

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
