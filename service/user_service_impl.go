package service

import (
	"context"
	"crypto/rand"
	"database/sql"
	"fmt"
	"math/big"
	"net/smtp"
	"time"
	"unicode"

	"survey/helper"
	"survey/model/domain"
	"survey/model/web"
	"survey/repository"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
	Error          error
	SMTPAuth       smtp.Auth
	SMTPHost       string
	SMTPPort       string
}

type Claims struct {
	NIM    string
	Email  string
	UserID int
	Role   string
	jwt.RegisteredClaims
}

type ClaimsPublic struct {
	Email  string
	UserID int
	Role   string
	jwt.RegisteredClaims
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate, smtpAuth smtp.Auth, smtpHost, smtpPort string) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
		SMTPAuth:       smtpAuth,
		SMTPHost:       smtpHost,
		SMTPPort:       smtpPort,
	}
}

func (service *UserServiceImpl) Register(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		fmt.Println("Validation error:", err)
		return web.UserResponse{} // or return an error response
	}

	tx, err := service.DB.Begin()
	if err != nil {
		fmt.Println("DB Begin error:", err)
		return web.UserResponse{} // or return an error response
	}
	defer helper.CommitOrRollback(tx)

	hashedPassword, err := HashPassword(request.Password)
	if err != nil {
		fmt.Println("HashPassword error:", err)
		return web.UserResponse{} // or return an error response
	}

	user := domain.User{
		NIM:      request.NIM,
		Email:    request.Email,
		Name:     request.Name,
		Password: hashedPassword,
		Role:     request.Role,
	}

	user = service.UserRepository.Register(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Login(ctx context.Context, request web.UserLoginRequest) (web.UserResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.UserResponse{}, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return web.UserResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByUsername(ctx, tx, request.NIM, request.Email)
	if err != nil {
		return web.UserResponse{}, err
	}

	err = ComparePassword(user.Password, request.Password)
	if err != nil {
		return web.UserResponse{}, err // Passwords don't match
	}

	token, err := GenerateToken(user.NIM, user.Email, user.Role, user.Id, "yourSecretKey")
	if err != nil {
		return web.UserResponse{}, err
	}

	userResponse := helper.ToUserResponse(user)
	userResponse.Token = token

	return userResponse, nil
}
func (service *UserServiceImpl) LoginPublic(ctx context.Context, request web.UserLoginPublicRequest) (web.UserResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.UserResponse{}, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return web.UserResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByUsernamePublic(ctx, tx, request.Email)
	if err != nil {
		return web.UserResponse{}, err
	}

	err = ComparePassword(user.Password, request.Password)
	if err != nil {
		return web.UserResponse{}, err // Passwords don't match
	}

	token, err := TokenUserPublic(user.Email, user.Role, user.Id, "yourSecretKey")
	if err != nil {
		return web.UserResponse{}, err
	}

	userResponse := helper.ToUserResponse(user)
	userResponse.Token = token

	return userResponse, nil
}

func (service *UserServiceImpl) SendResetPassword(ctx context.Context, request web.ForgotPasswordRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.UserResponse{Error: "Invalid request"}
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return web.UserResponse{Error: "Database error"}
	}
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByUsernamePublic(ctx, tx, request.Email)
	if err != nil {
		return web.UserResponse{Error: "User not found"}
	}

	otp := GenerateOTP()

	resetPassword := domain.ResetPassword{
		UserId:     user.Id,
		Token:      otp,
		Expired_at: time.Now().Add(2 * time.Minute),
	}

	resetPassword, err = service.UserRepository.InsertResetPassword(ctx, tx, resetPassword)
	if err != nil {
		return web.UserResponse{Error: "Failed to insert reset password"}
	}

	err = SendOTPResetPassword(user.Email, otp, service.SMTPHost, service.SMTPPort, service.SMTPAuth)
	if err != nil {
		return web.UserResponse{Error: "Failed to send OTP"}
	}

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) VerifyResetPassword(ctx context.Context, request web.ResetPasswordRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.UserResponse{Error: "Invalid request"}
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return web.UserResponse{Error: "Database error"}
	}
	defer helper.CommitOrRollback(tx)

	resetPassword, err := service.UserRepository.FindByToken(ctx, tx, request.Token)
	if err != nil {
		// Log the error for debugging purposes
		fmt.Printf("Error finding reset password by token: %v\n", err)
		return web.UserResponse{Error: "Invalid token"}
	}

	if time.Now().After(resetPassword.Expired_at) {
		return web.UserResponse{Error: "Token expired"}
	}

	user, err := service.UserRepository.FindById(ctx, tx, resetPassword.UserId)
	if err != nil {
		// Log the error for debugging purposes
		fmt.Printf("Error finding user by ID: %v\n", err)
		return web.UserResponse{Error: "User not found"}
	}

	// Here we assume GenerateToken takes only user ID and role for simplicity
	token, err := GenerateToken(user.NIM, user.Email, user.Role, user.Id, "yourSecretKey")
	if err != nil {
		// Log the error for debugging purposes
		fmt.Printf("Error generating token: %v\n", err)
		return web.UserResponse{Error: "Failed to generate token"}
	}

	userResponse := helper.ToUserResponse(user)
	userResponse.Token = token

	return userResponse
}

func (service *UserServiceImpl) ResetPassword(ctx context.Context, request web.ChangePasswordRequest, userId int) web.UserResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.UserResponse{Error: "Invalid request"}
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return web.UserResponse{Error: "Database error"}
	}
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		return web.UserResponse{Error: "User not found"}
	}

	if !IsValidPassword(request.Password) {
		return web.UserResponse{Error: "Password must be at least 8 characters long and include an uppercase letter, a lowercase letter, a number, and a special character"}
	}

	hashedPassword, err := HashPassword(request.Password)
	if err != nil {
		return web.UserResponse{Error: "Failed to hash password"}
	}

	user.Password = hashedPassword

	user, err = service.UserRepository.UpdatePassword(ctx, tx, user)
	if err != nil {
		return web.UserResponse{Error: "Failed to update password"}
	}

	user, err = service.UserRepository.DeletedByUserId(ctx, tx, userId)
	if err != nil {
		return web.UserResponse{Error: "Failed to reset password"}
	}

	return helper.ToUserResponse(user)
}

// generate token with claims username and role
func GenerateToken(nim string, email string, role string, userId int, secretKey string) (string, error) {
	// Set custom claims
	claims := &Claims{
		NIM:    nim,
		Email:  email,
		UserID: userId,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString([]byte("secretKey"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func IsValidPassword(password string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	// Panjang minimal 8 karakter
	if len(password) >= 8 {
		hasMinLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	// Setidaknya harus ada satu karakter dari setiap kriteria
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

func TokenUserPublic(email string, role string, userId int, secretKey string) (string, error) {
	// Set custom claims
	claims := &ClaimsPublic{
		Email:  email,
		UserID: userId,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString([]byte("secretKey"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func GenerateOTP() string {
	const number = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const otpLength = 6

	otp := make([]byte, otpLength)
	for i := range otp {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(number))))
		if err != nil {
			// Handle error
			return ""
		}
		otp[i] = number[randomIndex.Int64()]
	}
	return string(otp)
}

func SendOTPResetPassword(email, otp, host, port string, auth smtp.Auth) error {
	from := ""
	msg := fmt.Sprintf("To: %s\r\nSubject: Reset Password\r\n\r\nYour Code for reset password is: %s\r\n", email, otp)
	return smtp.SendMail(host+":"+port, auth, from, []string{email}, []byte(msg))
}
