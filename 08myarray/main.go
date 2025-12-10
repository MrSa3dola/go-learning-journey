package main

import "fmt"

func main() {
	fmt.Println("welcome to arrays\n")
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Banana"
	fmt.Println("fruits are", fruitList)

	fmt.Println("fruits are", len(fruitList)) // reserved memory (fixed)

	var vegList = []string{"potato", "beans", "mashroom"}
	fmt.Println("veg list", vegList)
	fmt.Println("veg list", len(vegList))
}
