package server

import (
	"Go_Sample_Server/handler"
	server "Go_Sample_Server/server/functions"
	"Go_Sample_Server/structures"
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
)

type Server struct {}
type FO = structures.FO;
type layoutMasterSet = structures.LayoutMasterSet
type simplePageMaster = structures.SimplePageMaster
type regionBody = structures.RegionBody
type pageSequence = structures.PageSequence
type flow = structures.Flow
type variableTextLayout = structures.VariableTextLayout
type variableText = structures.VariableText

func (h Server) GetVersion(ctx echo.Context) error {


	fo := &FO{XmlnsFo: "http://www.w3.org/1999/XSL/Format"}
	LMS := &layoutMasterSet{}
	SPM := &simplePageMaster{Page_width: "210mm", Page_height: "297mm", Master_name:  "test"}
	PS := &pageSequence{MasterReference: "test"}
	flow := &flow{FlowName: "xsl-region-body" }
	VTE := &variableTextLayout{AbsolutePosition: "absolute", Top: "0mm", Left: "0mm", Width: "200mm",Height: "50mm",  Overflow: "overflow", ZIndex: "0", FontSize: "30"}
	VT := &variableText{WrapOption: "wrap", LineHeight: "1.16", StartIndent: "0mm", FontFamily: "IPAMincho", Value: "123-0226"}
	
	VTE.VariableText = append(VTE.VariableText, *VT)
	flow.VariableTextLayout = append(flow.VariableTextLayout, *VTE)
	PS.Flow = append(PS.Flow, *flow)
	SPM.RegionBody = append(SPM.RegionBody, regionBody{Margin: "0"})
	LMS.SimplePageMaster = append(LMS.SimplePageMaster, *SPM)
	fo.MasterSet = append(fo.MasterSet, *LMS)
	fo.PageSequence = append(fo.PageSequence, *PS)
	output, err := xml.MarshalIndent(fo, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	server.WriteFoFile(string(output))

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