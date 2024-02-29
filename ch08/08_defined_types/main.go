package main

import "fmt"

type part struct {
	description string
	count       int
}

type car struct {
	name     string
	topSpeed float64
}

func main() {
	porsche := car{
		name:     "Porsche 911 R",
		topSpeed: 323.0,
	}
	fmt.Println("Name: ", porsche.name)
	fmt.Println("Top Speed: ", porsche.topSpeed)

	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 34
	fmt.Println("Description: ", bolts.description)
	fmt.Println("Count: ", bolts.count)
}
