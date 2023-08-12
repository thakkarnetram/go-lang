package main

import "fmt"

func main() {
	fmt.Println("Control flow w If else in Go")
	count := 5
	var result string

	if count < 5 {
		result="Less than 5"
	}else if count > 10 {
		result="Greater than 10"
	}else {
		result="Equal"
	}
	fmt.Println(result)
	
	if 10%2==0{
		fmt.Println("Even")
	}else{
		fmt.Println("Odd")
	}

	if num:=10;num>10{
		fmt.Println("Number greater than 10")
	}else {
		fmt.Println("Less than 10")
	}
}