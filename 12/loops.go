package main

import "fmt"

func main() {
	fmt.Println("Loops in GO")
	fmt.Println("")
	// there is only For Loop in GO and it has different variation . ! 
	days:=[]string{"Sunday","Saturday","Wednesday","Friday","Monday"}
	// prefer slices instead of arrays always 
	// variation 1
	for i:=0; i<len(days);i++{ // ++i , --i N/A 
		fmt.Println("Days -> ", days[i])
	}
	fmt.Println("")
	lang:=[]string{"R","C++","GO","Java","Python","Rust"}
	// variation 2
	for x:= range lang{
		fmt.Println("Languages Index -> " , x) // if i do this i get the index of the values
		fmt.Println("")
		fmt.Println("Languages Value -> " , lang[x]) // if i do this i get the values
	}
	fmt.Println("")
	marks:=[]int{33,57,88,10}
	// variation 3
	for index,mrks := range marks {
		fmt.Printf("Index of marks is %v and the marks are %v \n", index , mrks)
	}
	fmt.Println("")
	// variation 4
	for _ , mrks := range marks {
		fmt.Printf("The marks are %v \n", mrks)
	}
	// other things break, continue and goto 
	value:= 10
	for value < 20{
		if  value == 11 {
			goto label
		}
		if value == 12 {
			value++
			continue
		}
		if value == 15 {
			break
		}
		fmt.Println("Value -> " , value)
		value+=1
	}
	// goto statement 
	label:
	fmt.Println("Im the label xD")
}