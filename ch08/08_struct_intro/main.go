package main

import (
	"fmt"
)

func main() {
	var myStruct struct {
		number float64
		word   string
		toggle bool
	}
	fmt.Printf("%#v\n", myStruct)
	myStruct.toggle = true
	fmt.Printf("%#v\n", myStruct)
	myStruct.number = 56.3
	fmt.Printf("%#v\n", myStruct)
	myStruct.word = "Hello"
	fmt.Println(myStruct.number)
	fmt.Println(myStruct.word)
	fmt.Println(myStruct.toggle)
}
