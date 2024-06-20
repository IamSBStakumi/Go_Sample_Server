package main

import (
	"context"
	"log"
	"net/http"

	"Go_Sample_Server/handler"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"google.golang.org/api/option"
)

type firebaseApp struct {
	*firebase.App
}

func InitFirebaseApp() (*firebaseApp, error){
	opt := option.WithCredentialsFile("./firebase-adminsdk.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return &firebaseApp{app}, nil
}

func (app *firebaseApp) InitAuthService()(*auth.Client, error){
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth Client: %v\n",err)
	}

	return client, nil
}

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
		// TODO: Frontのアドレスに変更する
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))

	server := Server{}

	app, _ := InitFirebaseApp()
	client, _ := app.InitAuthService()

	handler.RegisterHandlers(e, server, *client)

	e.Logger.Fatal(e.Start("localhost:9000"))
}