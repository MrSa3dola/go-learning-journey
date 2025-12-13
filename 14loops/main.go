package main

import "fmt"

func main() {
	fmt.Println("welcome to loops")
	days := []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for _, day := range days {
		fmt.Printf("the day is %v\n", day)
	}

	roVal := 1
	for roVal < 10 { // similar to while loop
		fmt.Println("value is", roVal)
		if roVal == 2 {
			goto saad
		}
		if roVal == 5 {
			break
		}
		roVal++
	}
saad:
	fmt.Println("jumping at saad")
}
