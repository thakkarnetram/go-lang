package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch cases")
	rand.Seed(time.Now().UnixNano())
	diceNum:=rand.Intn(6)+1
	switch diceNum {
	case 1:
		fmt.Println("Dice val is 1")
	case 2:
		fmt.Println("Dice val is 2")
		// fallthrough  when it hits the Case 2 , it also will hit the Case 3
	case 3:
		fmt.Println("Dice val is 3")
	case 4:
		fmt.Println("Dice val is 4")
	case 5:
		fmt.Println("Dice val is 5")
	case 6:
		fmt.Println("Dice val is 6")
	default:
		fmt.Println("Default Dice val is ", diceNum )
	}
}