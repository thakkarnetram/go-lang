package main

import "fmt"

func main() {
	fmt.Println("Pointers in GO ")
	var ptr *int 
	fmt.Println("Ptr ", ptr)
	number:= 20
	var pointer = &number
	fmt.Println("Pointer address &  - > " , pointer )
	fmt.Println("Pointer  value * - > " , *pointer )
	*pointer = *pointer - 10
	fmt.Println("New value -> " , number)
	fmt.Println("New value -> " , &number)
}