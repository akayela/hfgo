package main

import (
	"fmt"
)

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func main() {
	subscriber1 := subscriber{
		name:   "Audience Kayela",
		rate:   12.99,
		active: true,
	}
	fmt.Println("Name: ", subscriber1.name)

	var subscriber2 subscriber
	subscriber2.name = "Joana Kayela"
	subscriber2.rate = 12.99
	subscriber2.active = true
	fmt.Println("Name: ", subscriber2.name)

}
