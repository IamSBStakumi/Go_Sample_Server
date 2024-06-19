package main

import (
	"net/http"

	"Go_Sample_Server/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {}

func (h Server) GetVersion(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "0.0.1")
}

func (h Server) RegisterUser(ctx echo.Context) error {
	return ctx.JSON(http.StatusCreated, "Success!")
}

func (handle Server) DeleteUser(ctx echo.Context, firebaseUid string) error{
	return ctx.JSON(http.StatusOK, "User deleted")
}

func main(){
	e := echo.New()

	// middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))

	server := Server{}


	handler.RegisterHandlers(e, server)

	e.Logger.Fatal(e.Start("localhost:9000"))
}