package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// create file OS package
	// reading n other stuff using Ioutil package
	fmt.Println("Files in Go")
	content:="he"
	file,err:=os.Create("./file.txt")
	checkNilError(err) // panic keyword stops the execution of func and shows the err 
	len,err:=io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("length of file  ", len)
	defer file.Close()
	readFile("./file.txt")
}

func readFile(file string)  {
	// not read in string format
	// read in bytes
	data,err:=ioutil.ReadFile(file)
	checkNilError(err)
	fmt.Println("data " ,string(data))
}

func checkNilError(err error){
	if err != nil{
		panic(err)
	}
}