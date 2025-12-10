package main

import "fmt"

func main() {
	// structs
	fmt.Println("structs")
	// no inheritance, no parent-child relationship
	saad := User{"Saad", "saad@ga.com", true, 20}
	fmt.Println(saad)
	fmt.Printf("details are %+v\n", saad)
	fmt.Printf("Name is %v\n", saad.Name)
	fmt.Printf("Email is %v\n", saad.Email)
}

type User struct { // capital U makes it public
	Name   string
	Email  string
	Status bool
	Age    int
}
