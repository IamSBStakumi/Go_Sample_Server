package server

import (
	"Go_Sample_Server/handler"
	"context"
	"log"
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
)

type Server struct {}

func (h Server) GetVersion(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "0.0.1")
}

func (h Server) RegisterUser(ctx echo.Context, client *auth.Client) error {
	var req handler.RegisterUserJSONRequestBody
	if err := ctx.Bind(&req);err !=nil {
		log.Fatalf("request body is not Found: %v\n", err)
	}

	params := (&auth.UserToCreate{}).
		Email(string(req.Email)).
		EmailVerified(false).
		Password(req.Password).
		DisplayName(req.Username).
		Disabled(false)

	_, err := client.CreateUser(context.Background(), params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}

	return ctx.JSON(http.StatusCreated, "Success!")
}

func (handle Server) DeleteUser(ctx echo.Context, firebaseUid string, client *auth.Client) error{
	err := client.DeleteUser(context.Background(), firebaseUid)
	if err != nil {
		log.Fatalf("error deleting user: %v\n", err)
	}
	
	return ctx.JSON(http.StatusOK, "User deleted")
}