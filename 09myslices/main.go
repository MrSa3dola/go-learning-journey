package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("slices")

	var fruitList = []string{"apple", "banana", "grape"}

	fmt.Printf("type is %T\n", fruitList)

	fruitList = append(fruitList, "Mango")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3]) // 3 not included [1,2] only
	fmt.Println(fruitList)

	highScore := make([]int, 4) // type, size
	highScore[0] = 234
	highScore[1] = 54
	highScore[2] = 25
	highScore[3] = 76
	fmt.Println(highScore)
	fmt.Println(len(highScore))
	highScore = append(highScore, 555, 666, 777) // reallocate the momory and save the new values
	fmt.Println(highScore)
	fmt.Println(len(highScore))
	fmt.Println("--------------")
	fmt.Println(sort.IntsAreSorted(highScore)) // is sorted?
	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println("=============")
	// how to remove a value from slices based on index
	var courses = []string{"react", "java", "swift", "python", "ruby"}
	fmt.Println(courses)

	var idx int = 2
	courses = append(courses[:idx], courses[idx+1:]...) // remove the value at index 2
	// append is more optimized
	fmt.Println(courses)
}
