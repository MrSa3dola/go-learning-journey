package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// switch-case
	fmt.Println("switch-case")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Dice number is", diceNumber)
	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spot")
		fallthrough // if it got 3, it will also execute case 4 and 5
	case 4:
		fmt.Println("You can move 4 spot")
		fallthrough // if it got 4, it will also execute case 5
	case 5:
		fmt.Println("You can move 5 spot")
	case 6:
		fmt.Println("Dice value is 6 and you can open and roll the dice again")
	default:
		fmt.Println("Invalid dice number")
	}
}
