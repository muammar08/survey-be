package web

type UserLoginRequest struct {
	NIM      string `json:"nim" validate:"max=200,min=0"`
	Email    string `json:"email" validate:"max=200,min=0"`
	Password string `json:"password" validate:"required,max=200,min=1"`
}

type UserLoginPublicRequest struct {
	Email    string `json:"email" validate:"max=200,min=1"`
	Password string `json:"password" validate:"required,max=200,min=1"`
}
