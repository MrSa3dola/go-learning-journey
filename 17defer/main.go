package main

import "fmt"

func main() {
	// defer technique LIFO: last in first out
	defer fmt.Println("hello") // it will be run at the end of the main braces
	fmt.Println("world")
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	myDefer()
}

func myDefer(){
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

