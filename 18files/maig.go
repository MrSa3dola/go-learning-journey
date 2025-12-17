package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("files")
	content := "This needs to go in a file - saad.in"
	
	file, err := os.Create("./msaadfile.txt")
	checkNilErr(err)
	len, err := io.WriteString(file, content)

	checkNilErr(err)
	fmt.Println("length is", len)
	defer file.Close()
	readFile("./msaadfile.txt")
}		

func readFile(fileName string){
	dataByte, err := ioutil.ReadFile(fileName)
	checkNilErr(err)
	fmt.Println("text is\n", string(dataByte))
}

func checkNilErr(err error){
	if err != nil {
		panic(err)
	}
}