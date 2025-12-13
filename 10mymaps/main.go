package main

import "fmt"

func main() {
	fmt.Println("maps")
	languages := make(map[string]string)
	languages["js"] = "JavaScript"
	languages["py"] = "Python"
	languages["rb"] = "Ruby"
	fmt.Println("list of languages:", languages)
	fmt.Println("js shorts for", languages["js"])

	delete(languages, "rb")
	fmt.Println("list of languages:", languages)

	// loops
	for key, value := range languages {
		fmt.Printf("FOR key %v, value is %v\n", key, value)
	}
}
