package main

import (
	"fmt"
	"time"
)

func main() {
	go hello("Hi")
	hello("Niqa")
}

func hello(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3*time.Microsecond)
		fmt.Println(s)
	}
}