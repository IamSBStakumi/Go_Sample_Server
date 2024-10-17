package server

import (
	"Go_Sample_Server/handler"
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"

	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
)

type Server struct {}

type FO struct {
	XMLName xml.Name `xml:"fo:root"`
	MasterSet  []masterSet   `xml:"fo:layout-master-set"`
	XmlnsFo string   `xml:"xmlns:fo,attr"`
}

type masterSet struct {
	XMLName xml.Name `xml:"fo:simple-page-master"`
	RegionBody []regionBody
	Page_width string `xml:"page-width,attr"`
	Page_height string `xml:"page-height,attr"`
	Master_name string `xml:"master-name,attr"`
}

type regionBody struct {
	XMLName xml.Name `xml:"fo:region-body"`
	Margin string `xml:"margin,attr"`
	AutoClose []string
}

// foファイルの書き出し
func WriteFoFile(fo string){
	f, err := os.Create("./xml/sample.fo")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	defer f.Close()

	d := []byte(xml.Header + fo)

	n, err := f.Write(d)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Printf("%d bytes書き込みました。", n)
}

func (h Server) GetVersion(ctx echo.Context) error {
	fo := &FO{XmlnsFo: "http://www.w3.org/1999/XSL/Format"}
	MS := &masterSet{Page_width: "210mm", Page_height: "297mm", Master_name:  "test"}
	MS.RegionBody = append(MS.RegionBody, regionBody{Margin: "0"})
	fo.MasterSet = append(fo.MasterSet, *MS)
	output, err := xml.MarshalIndent(fo, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	WriteFoFile(string(output))

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