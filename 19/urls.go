package main

import (
	"fmt"
	"net/url"
)

const myUri string = "https://www.omdbapi.com/?i=tt0167260&apikey=f7292e1c"

func main() {
	fmt.Println("Handling urls in GO")

	// parse
	res,_:=url.Parse(myUri)
	fmt.Println("Res -> " , res.Scheme)
	fmt.Println("Res -> " , res.Path)
	
	qparam:=res.Query()
	fmt.Printf("qparam type %T \n ",qparam)
	fmt.Print("qparam values \n",qparam["apikey"]) // extract
	fmt.Println("qparam values \n",qparam)

	// use for loop
	for _,val := range qparam {
		fmt.Println("Params values are  : ",  val)
	}
	
	// parts of url
	// always pass referance & 
	partsUrl:= &url.URL{
		Scheme: "https",
		Host: "www.google.com",
		Path: "/search",
		RawQuery:"q=ok",
	}

	// construct
	constructUrl:= partsUrl.String()
	fmt.Println(constructUrl)
}