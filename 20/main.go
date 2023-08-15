package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	GetReq()
	PostReq()
	FormReq()
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

func PostReq(){
	const url  = "http://localhost:1111/post"
	// data payload
	reqBody := strings.NewReader(`
		{
			"name":"Devilop",
			"role":"backendDev"
		}
	`)
	res,err := http.Post(url,"application/json",reqBody)
	ErrorHandler(err)
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	ErrorHandler(err)
	fmt.Println("Data -> " , string(data))
}

func FormReq(){
	const myUrl  = "http://localhost:1111/postform"
	// adding form data 
	data:= url.Values{}
	data.Add("name","unkown")
	data.Add("age","19")
	data.Add("location","Mumbai")
	// send req
	res,err:= http.PostForm(myUrl,data)
	ErrorHandler(err)
	defer res.Body.Close()
	content,err:=io.ReadAll(res.Body)
	ErrorHandler(err)
	fmt.Println("Form content ", string(content))

}

func ErrorHandler(err error){
	if err != nil{
		panic(err)
	}
}