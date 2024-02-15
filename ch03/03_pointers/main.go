package main

import (
	"fmt"
	"reflect"
)

func main() {
	// variable
	amount := 6

	// pointer to variable
	amountLoc := &amount

	// dereferencing pointer
	fmt.Println(*amountLoc)

	// Printing the memory addrss of the variable
	fmt.Println(amountLoc)

	// type of pointer
	var myInt int
	fmt.Println(reflect.TypeOf(&myInt))

	// type of pointer
	var myFloat float64
	fmt.Println(reflect.TypeOf(&myFloat))

	// type of pointer
	var myBool bool
	fmt.Println(reflect.TypeOf(&myBool))

	// variables that hold pointers
	var myInt2 int
	myIntPointer := &myInt2
	fmt.Println(myIntPointer)

	// getting or changing the value at a pointer
	myString := "Heeey!"

	myStringPointer := &myString

	valueAtPointer := *myStringPointer

	fmt.Println(valueAtPointer)

	fmt.Println("Changing the value at pointer")

	*myStringPointer = "Hello!"

	fmt.Println(*myStringPointer)
}
