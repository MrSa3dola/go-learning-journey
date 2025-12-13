package main

import "fmt"

func main() {
	fmt.Println("welcome to functions")

	greet()

	res := adder(3, 4)
	fmt.Println("result from adder is:", res)
}

func adder(x int, y int) int { // function signatures (params, return types)
	return x + y
}

func greet() {
	fmt.Println("hello from greet function")
}
