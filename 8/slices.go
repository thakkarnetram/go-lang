package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in GO")
	var list = []string{"Apple","Motorola","Samsung"}
	fmt.Println("List -> " , list)
	fmt.Printf("Type of List -> %T \n " , list)
	list = append(list, "Xiomi" , "Realme")	
	fmt.Println("Whole list ==> " , list)
	list =append(list[2:2]) // : used to slice the slice 
	// basically it slices like 
	// slice -> 0 , 1 , 2 , 3 , 4  
	// if its 1: the first element will be removed 
	// similary  when 1:2 
	fmt.Println("Sliced List 0 , 1  -> " , list)
	list =append(list[1:2])
	// the last range is non inclusive ie index 2 wont be counted
	fmt.Println("Sliced List New -> " , list)
	list=append(list[:2])
	fmt.Println("Sliced List New :2 -> " , list)

	// other way to initialize slice
	highScore := make([]int,2)
	highScore[0] = 222222
	highScore[1] = 22222
	// highScore[2] = 222 out of bound it will go cause there is no memory 
	// highScore[3] = 2222 out of bound it will go cause there is no memory 
	// but if we use append it will reallocate the memory according to the data input 
	highScore =  append(highScore,222, 22222,22222222,2222222222)
	fmt.Println("Highscore -> " ,  highScore)
	// sort package 
	fmt.Println(sort.IntsAreSorted(highScore))
	sort.Ints(highScore)
	fmt.Println("Highscore sorted -> " ,  highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	// removing value based on index
	var food = []string{
		"Whey Protein,",
		"Soya,",
		"Paneer,",
		"Eggs.",
	}
	fmt.Println("Food -> " , food)
	// use append to remove values
	index:=1 //  basically  from index 2  onwards 
	food =append(food[:index],food[index+1:]...) // add this data again 
	fmt.Println("Food Updated -> " , food)

}