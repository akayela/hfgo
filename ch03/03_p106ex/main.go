package main

import "fmt"

func main() {
	var myInt int

	myIntPointer := &myInt

	myInt = 42

	fmt.Println(*myIntPointer)
}
