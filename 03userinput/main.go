package main

import (
	"bufio"
	"fmt"
	"os"
)


func main(){
	welcome := "User input"
	fmt.Println(welcome)
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rate")

	// comman ok || err ok syntax
	input, err := reader.ReadString('\n')
	fmt.Println("The rate is", input)
	fmt.Printf("type of rate is %T \n", input)

	fmt.Println()

	fmt.Println(err) // <nil>
	
}