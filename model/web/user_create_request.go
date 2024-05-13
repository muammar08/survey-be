package web

type UserCreateRequest struct {
	NIM      string `json:"nim" validate:"max=200,min=0"`
	Email    string `json:"email" validate:"required,max=200,min=1"`
	Name     string `json:"name" validate:"required,max=200,min=1"`
	Password string `json:"password" validate:"required,max=200,min=1"`
	Role     string `json:"role"`
}
