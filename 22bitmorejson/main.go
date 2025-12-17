package main

import (
	"encoding/json"
	"fmt"
)

// keys better to be lowerCase
type course struct {
	Name     string `json:"coursename"` // alias as a coursename
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // don't pass this field
	Tags     []string `json:"tags,omitempty"` // don't pass the nil
}

func main() {
	fmt.Println("JSON")
	// EncodeJSON()
	DecodeJSON()
}

func EncodeJSON() {

	lcoCourses := []course{
		{
			"React JS bootcamp", 20, "learn.in", "12345", []string{"web-dev", "js"},
		},
		{
			"MERN bootcamp", 30, "learn.in", "abcd", []string{"full-stack", "js"},
		},
		{
			"Angular bootcamp", 40, "learn.in", "123gg", nil,
		},
	}
	// package this data as a JSON data
	// Marshal -> passing interface

	// finalJSON, err := json.Marshal(lcoCourses)
	// checkNil(err)
	// fmt.Printf("%s\n", finalJSON)

	//  if the tags slice is empty, it will return null not nil
	finalJSON, err := json.MarshalIndent(lcoCourses, "", "\t")
	checkNil(err)
	fmt.Printf("%s\n", finalJSON)
}
func DecodeJSON() {
	jsonData := []byte(`
	{
		"coursename": "React JS bootcamp",
		"Price": 20,
		"website": "learn.in",
		"tags": ["web-dev","js"]
	}
	`)
	// check if valid JSON data or not
	var lcoCourse course
	checkValid := json.Valid(jsonData)
	if checkValid {
		fmt.Println("JSON data is valid")
		json.Unmarshal(jsonData, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON not valid")
	}
	// some cases where you just want to add data to key value pair
	// interface type is defined as a set of method signatures
	/*
		type Abser interface {
			Abs() float64
		}
	*/
	// maybe it's int or string or slice, so we use interface
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonData, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}
	// the Price converted into float64 (defualt)
}
func checkNil(err error) {
	if err != nil {
		panic(err)
	}
}
