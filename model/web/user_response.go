package web

type UserResponse struct {
	Id       int     `json:"id"`
	NIM      *string `json:"nim"`
	Email    string  `json:"email"`
	Name     string  `json:"name"`
	Password string  `json:"password"`
	Role     string  `json:"role"`
	Token    string  `json:"token"`
	Error    string  `json:"error"`
}

type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
