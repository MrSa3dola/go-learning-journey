package main

import "fmt"

// jwtToken := 300000 // can't use walrus operator outside the methods, u should follow (var x = ..)

const LoginToken string = "gskljnglks"

// putting 'L' as a capital char to make it public

func main(){
	var username string = "Saad"
	fmt.Println(username);
	fmt.Printf("Variable is of type: %T \n", username);
	
	fmt.Println();

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn);
	fmt.Printf("Variable is of type: %T \n", isLoggedIn);
	
	fmt.Println();

	var smallVal uint16 = 255
	fmt.Println(smallVal);
	fmt.Printf("Variable is of type: %T \n", smallVal);

	fmt.Println();

	var smallFloat float32 = 426.56463164215 // 5 digits after the decimal point
	fmt.Println(smallFloat);
	fmt.Printf("Variable is of type: %T \n", smallFloat);

	fmt.Println();

	var bigFloat float64 = 426.5646316421556654545 // 13 digits after the decimal point
	fmt.Println(bigFloat)
	fmt.Printf("Variable is of type: %T \n", bigFloat);

	fmt.Println();

	var x int
	fmt.Println(x)
	fmt.Printf("Variable is of type: %T \n", x)

	fmt.Println();

	// implicit type
	var website = "code"
	fmt.Println(website);
	fmt.Printf("Variable is of type: %T \n", website)

	fmt.Println();

	// no var style
	num := 5000
	fmt.Println(num)

	fmt.Println();

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}	
