package main

import "fmt"

func main() {
	fmt.Println("Maps in Go")
	languages:=make(map[string]string)	
	languages["GO"] = "GoLang"
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	fmt.Println("Map -> " , languages)
	fmt.Printf("Map Type  ->  %T \n " , languages)

	//specific item
	fmt.Println("Map item -> " , languages["JS"])

	// delete item
	delete(languages,"PY") // doesnt throww err if key isnt there 
	fmt.Println("Map updated -> " , languages)

	// loops iteration
	for key , value := range languages {
		fmt.Printf(" For Key %v and value is %v \n " , key, value)  // todo read bout placeholders from doc 
	}
	// comma , ok syntax in for loop 
	for _ , value := range languages {
		fmt.Printf("value is %v \n " , value)  
	}
}