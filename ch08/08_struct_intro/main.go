package main

import (
	"fmt"
)

var myStruct struct {
	number float64
	word   string
	toggle bool
}

var subscriber struct {
	name   string
	rate   float64
	active bool
}

func main() {

	fmt.Printf("%#v\n", myStruct)
	myStruct.toggle = true
	fmt.Printf("%#v\n", myStruct)
	myStruct.number = 56.3
	fmt.Printf("%#v\n", myStruct)
	myStruct.word = "Hello"
	fmt.Println(myStruct.number)
	fmt.Println(myStruct.word)
	fmt.Println(myStruct.toggle)

	subscriber.name = "Audience"
	subscriber.rate = 4.99
	subscriber.active = true
	fmt.Println(subscriber.name)
	fmt.Println(subscriber.rate)
	fmt.Println(subscriber.active)
}
