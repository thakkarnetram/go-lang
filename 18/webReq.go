package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	res,err := http.Get(url)
	errHandler(err)
	fmt.Println("Res ",res )
	defer res.Body.Close() // always close the connection at the last

	// reading the res
	data,err:=io.ReadAll(res.Body)
	errHandler(err)
	content:=string(data)
	fmt.Println(content)
}

func errHandler(err error)  {
	if err != nil{
		panic(err)
	}
}