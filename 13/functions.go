package main

import "fmt"

func main() {
	fmt.Println("Functions in Go")
	result:=add(323,22)
	fmt.Println("Result -> " , result)
	fmt.Println("")
	proRes,msg:=pro(4,556,212)
	fmt.Println("Pro add -> ", proRes)
	fmt.Println("Pro Msg -> ", msg)
}
// mentioning the type in functions is know as function signatures
// what type of value is expected as input and expected as return
func add (one int , two int ) int {  
	return one+two
}

// pro functions to recieve n number of values as input
func pro(number...int)(int,string){
	total:=0
	for _, value := range number{
		total+=value
	}
	return total,"hi"
}