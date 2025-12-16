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
	saad.GetStatus()
	saad.NewMail()
}

type User struct { // capital U makes it public
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}
func (u User) NewMail() {
	u.Email = "test@daa.co"
	fmt.Println("Email of this user is:", u.Email)
}
