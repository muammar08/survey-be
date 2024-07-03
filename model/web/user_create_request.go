package web

type UserCreateRequest struct {
	NIM      string `json:"nim" validate:"omitempty,numeric,max=15,min=0"`
	Email    string `json:"email" validate:"email,required"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required,min=6,max=20"`
	Role     string `json:"role"`
}
