package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Welcome to time")
	presentTime := time.Now()
	fmt.Println(presentTime)
	
	// month-day-year format use 01-02-2006 standard
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// month is time.Month and the rest are int
	createdDate := time.Date(2020, time.May, 1, 23,23,0,0,time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}