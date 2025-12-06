package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main(){
	fmt.Println("welcome to pizza app")
	fmt.Println("please rate our pizza between 1 : 5")

	reader := bufio.NewReader(os.Stdin)

	in, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating,",in)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(in), 64)
	// strings.TrimSpace to delete the '\n'
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to the rate:", numRating + 1)
	}
}