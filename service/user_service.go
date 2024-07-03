package service

import (
	"context"
	"survey/model/web"
)

type UserService interface {
	Register(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Login(ctx context.Context, request web.UserLoginRequest) (web.UserResponse, error)
	LoginPublic(ctx context.Context, request web.UserLoginPublicRequest) (web.UserResponse, error)
}
