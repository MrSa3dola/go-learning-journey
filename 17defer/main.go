package main

import "fmt"

func main() {
	// defer technique LIFO: last in first out
	defer fmt.Println("hello") // it will be run at the end of the main braces
	fmt.Println("world")

}
