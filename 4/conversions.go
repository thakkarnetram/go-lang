package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Go Conversions")
	fmt.Println("Rate GO lang between 1 - 10 ")

	rating := bufio.NewReader(os.Stdin)
	input, _ := rating.ReadString('\n')
	fmt.Println("Thanks for rating GO -> ", input)

	numRating, error := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if error != nil {
		fmt.Println(error)
	}else {
		fmt.Println("minus -1 from your rating " , numRating-1)
	}
}
