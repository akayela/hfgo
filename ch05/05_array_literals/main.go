package main

import "fmt"

func main() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	primes := [5]int{2, 3, 5, 7, 11}
	text := [3]string{
		"This is a series of long strings",
		"which would be awkward to place",
		"together on a single line",
	}
	fmt.Println(text, notes, primes)

}
