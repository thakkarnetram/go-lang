package main

import "fmt"

func main() {
	defer fmt.Println("Defer statement") // when defer is used the line is executed at last 
	// suppose there are 3 defer statement
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	// in defer the concept is First In Last Out 
	// executed like Three ->  Two -> One -> Defer statement
	fmt.Println("Non defer statement")
	myDefer2()
	myDefer()
}

func myDefer()  {
	for i:=0;i<5;i++{
		defer fmt.Println("Defer 1 -> " , i)
	}
}

func myDefer2 (){
	for x:=0;x<=5;x++{
		defer fmt.Println("Defer 2 -> " , x)
	}
}