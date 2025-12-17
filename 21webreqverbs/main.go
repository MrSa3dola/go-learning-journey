package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("web verb video - LCO")
	url := "http://localhost:8000"
	PerformGetRequest(url)
}

func PerformGetRequest(url string) {
	response, err := http.Get(url)
	checkNil(err)
	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length", response.ContentLength)
	
	// string builder is more effecient  
	var responseString strings.Builder
	content, err  := io.ReadAll(response.Body)
	checkNil(err)
	byteCount, _ := responseString.Write(content)
	fmt.Println("byte count is", byteCount) // length
	fmt.Println("raw data is:", responseString.String())

	// fmt.Println(content)
	// fmt.Println(string(content))
}

func checkNil(err error){
	if err != nil{
		panic(err)
	}
}