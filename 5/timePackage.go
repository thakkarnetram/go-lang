package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time package")
	currentTime := time.Now();
	fmt.Println("Current time " , currentTime)
	fmt.Println("formated time ", currentTime.Format("01-02-2006 15:04:05 Monday"))
	createdDate := time.Date(2024,time.September,10,10,0,0,0,time.UTC)
	fmt.Println("Created date ", createdDate)
}