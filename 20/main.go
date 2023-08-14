package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	GetReq()
}

func GetReq() {
	const url = "http://localhost:1111/get"

	res, err := http.Get(url)
	ErrorHandler(err)

	fmt.Println("Res status -> " , res.StatusCode)

	content,err:= io.ReadAll(res.Body)
	ErrorHandler(err)
	fmt.Println("Content -> " , string(content))
	// other way to read data
	var resString strings.Builder
	bytes,err:=resString.Write(content)
	ErrorHandler(err)
	fmt.Println("Byte count -> ", bytes)
	fmt.Println("Content the other way  -> ", resString.String())
	
	defer res.Body.Close()
}

func ErrorHandler(err error){
	if err != nil{
		panic(err)
	}
}