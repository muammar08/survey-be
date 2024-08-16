package web

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,max=200,min=1"`
}

type ResetPasswordRequest struct {
	Token string `json:"otp" validate:"required,max=200,min=1"`
}

type ChangePasswordRequest struct {
	Password string `json:"new_password" validate:"required,max=200,min=1"`
}
