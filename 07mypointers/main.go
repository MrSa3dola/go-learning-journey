package main

import "fmt"

func main(){
	fmt.Println("welcom to pointers")

	// var ptr *int // pointer for holding int value
	// fmt.Println("Value of pointer is", ptr)
	myNum := 23

	var ptr = &myNum // refrencing to some memory
	fmt.Println("address of actual pointer is", ptr)
	fmt.Println("value of actual pointer is", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is", myNum)
}
