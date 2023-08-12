package main

import "fmt"

func main() {
	fmt.Println("Arrays in GO")
	var animals[4] string
	animals[0] = "Dog"
	animals[2] = "Cat"
	animals[3] = "Mouse"
	fmt.Println("Array -> " , animals)
	fmt.Println("Array -> " , len(animals))

	var locations = [2] string{"Mumbai","Goa"}
	fmt.Println("Locations -> " , locations)
}