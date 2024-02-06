package main

import (
	"fmt"
)

func main() {
	// Single variable definition
	var customerName string = "Audience Kayela"

	// Single variable declaration
	// When you declare a variable without assignin a value,
	// it will be assigned the zero value of its type.
	var quantity int

	// Multiple variable declaration
	var length, width float64

	// Single variable initialisation
	quantity = 2

	// Multiple variable initialisation
	length, width = 1.2, 2.4

	fmt.Println(customerName)

	fmt.Println("Has orderded", quantity, "sheets")

	fmt.Println("each with an area of")

	fmt.Println(length*width, "square meters")
}
