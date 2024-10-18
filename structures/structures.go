package structures

import (
	"encoding/xml"
)

type FO struct {
	XMLName xml.Name `xml:"fo:root"`
	MasterSet  []LayoutMasterSet
	PageSequence []PageSequence
	XmlnsFo string   `xml:"xmlns:fo,attr"`
}

type LayoutMasterSet struct {
	XMLName xml.Name `xml:"fo:layout-master-set"`
	SimplePageMaster []SimplePageMaster
}

type SimplePageMaster struct {
	XMLName xml.Name `xml:"fo:simple-page-master"`
	RegionBody []RegionBody
	Page_width string `xml:"page-width,attr"`
	Page_height string `xml:"page-height,attr"`
	Master_name string `xml:"master-name,attr"`
}

type RegionBody struct {
	XMLName xml.Name `xml:"fo:region-body"`
	Margin string `xml:"margin,attr"`
	HTMLAutoClose []string
}

type PageSequence struct {
	XMLName xml.Name `xml:"fo:page-sequence"`
	MasterReference string `xml:"master-reference,attr"`
	Flow []Flow
}

type Flow struct {
	XMLName xml.Name `xml:"fo:flow"`
	FlowName string `xml:"flow-name,attr"`
	blockContainer []BlockContainer
	VariableTextLayout []VariableTextLayout
}

type BlockContainer struct {
	XMLName xml.Name `xml:"fo:block-container"`
	AbsolutePosition string `xml:"absolute-position,attr"`
	Top string `xml:"top,attr"`
	Left string `xml:"left,attr"`
	Width string `xml:"width,attr"`
	ZIndex string `xml:"z-index,attr"`
	VariableTextLayout []VariableTextLayout
}

type VariableTextLayout struct {
	XMLName xml.Name `xml:"fo:block-container"`
	AbsolutePosition string `xml:"absolute-position,attr"`
	Top string `xml:"top,attr"`
	Left string `xml:"left,attr"`
	Width string `xml:"width,attr"`
	Height string `xml:"height,attr"`
	ZIndex string `xml:"z-index,attr"`
	Overflow string `xml:"overflow,attr"`
	FontSize string `xml:"font-size,attr"`
	VariableText []VariableText
}

type VariableText struct {
   	XMLName xml.Name `xml:"fo:block"`
	WrapOption string `xml:"wrap-option,attr"`
	LineHeight string `xml:"line-height,attr"`
	StartIndent string `xml:"start-indent,attr"`
	FontFamily string `xml:"font-family,attr"`
	Value string `xml:",chardata"`
}
