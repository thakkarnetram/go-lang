package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := " Hey rate our code "
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println(" github.com/op -> ")

	// comma ok || err err operator  (just like try/catch )
	in , _ := reader.ReadString('\n');
	fmt.Println("Thank u " , in)
	billing()
}

func billing (){ 
	read := bufio.NewReader(os.Stdin)
	fmt.Print("Order Something  -> " )
	input , _ := read.ReadString('\n')
	fmt.Println("Your order is " , input)
}