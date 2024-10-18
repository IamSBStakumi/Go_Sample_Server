package server

import (
	"encoding/xml"
	"fmt"
	"os"
)

// foファイルの書き出し
func WriteFoFile(fo string){
	f, err := os.Create("./xml/sample_1.fo")
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