package main

import "fmt"

func main() {
	fmt.Println("Structs in GO ")
	// structs are like classes
	// but no inheritance, no super classes
	devil:=User{"Devil","ok@mail",16,"ojksaje"}
	fmt.Println("User -> " , devil)
	fmt.Println("User Email  -> " , devil.Email)
	fmt.Printf("User Type %T \n " , devil)
	// depth details
	fmt.Printf("User details %+v \n",devil)
}

type User struct {
	Name string
	Email string
	Age int
	Password string
}