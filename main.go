package main

import (
	"context"
	"log"
	"net/http"

	"Go_Sample_Server/handler"
	"Go_Sample_Server/server"

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

func main(){
	e := echo.New()

	// middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// TODO: Frontのアドレスに変更する
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))

	server := server.Server{}

	app, _ := InitFirebaseApp()
	client, _ := app.InitAuthService()

	handler.RegisterHandlers(e, server, *client)

	e.Logger.Fatal(e.Start("localhost:9000"))
}