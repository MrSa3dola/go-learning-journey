package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.google.com/"

func main() {
	fmt.Println("web requests")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("response is of type: %T\n", response)
	fmt.Println("------")
	// fmt.Println("response is \n", response)

	defer response.Body.Close() // caller's responsibility to close the connection

	dataBytes, err := io.ReadAll(response.Body)
	if err != nil{
		panic(err)
	}
	content := string(dataBytes)
	fmt.Println("the body is", content)
}
