package main

import "fmt"

func main() {
	// if-else
	fmt.Println("if-else")
	loginCount := 23
	var res string
	if loginCount < 10 {
		res = "regular user"
	} else {
		res = "new user"
	}
	fmt.Println(res)
	// web request handling
	if num := 33; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("Not less than 10")
	}
}
