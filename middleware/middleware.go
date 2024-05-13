package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/web"

	"github.com/golang-jwt/jwt/v4"
	"github.com/julienschmidt/httprouter"
)

type Middleware struct {
	Handler httprouter.Handle
}

func NewAuthMiddleware(handler httprouter.Handle) *Middleware {
	return &Middleware{
		Handler: handler,
	}
}

type Claims struct {
	Email  string
	UserID int
	Role   string
	jwt.RegisteredClaims
}

func (middleware *Middleware) ApplyMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		tokenString := request.Header.Get("Authorization")
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secretKey"), nil
		})

		if err != nil || !token.Valid || claims.Role == "ADMIN" {
			fmt.Println("Token invalid or parsing error:", err)
			helper.WriteResponse(writer, http.StatusUnauthorized, web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
			})
			return
		}

		ctx := context.WithValue(request.Context(), "email", claims.Email)
		ctx = context.WithValue(ctx, "userId", claims.UserID)
		ctx = context.WithValue(ctx, "role", claims.Role)
		request = request.WithContext(ctx)

		next(writer, request, params)
	}
}

func (middleware *Middleware) ApplyAdminMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		tokenString := request.Header.Get("Authorization")
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secretKey"), nil
		})

		if err != nil || !token.Valid || claims.Role == "USER" {
			fmt.Println("Token invalid or parsing error:", err)
			helper.WriteResponse(writer, http.StatusUnauthorized, web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
			})
			return
		}

		ctx := context.WithValue(request.Context(), "email", claims.Email)
		ctx = context.WithValue(ctx, "userId", claims.UserID)
		ctx = context.WithValue(ctx, "role", claims.Role)
		request = request.WithContext(ctx)

		next(writer, request, params)
	}
}
