package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Strings
	fmt.Println("Hello, \nGo!")
	fmt.Println("Hello, \tGo!")
	fmt.Println("Quotes, \"\"")
	fmt.Println("Backslash, \\")

	// Runes aka Char in other languages
	fmt.Print('A')
	fmt.Print('c')

	// Numbers
	fmt.Println(42)
	fmt.Println(1 + 6)

	// Comparison
	fmt.Println(4 < 6)
	fmt.Println(4 > 6)
	fmt.Println(4 == 6)
	fmt.Println(4 != 6)
	fmt.Println(4 <= 6)
	fmt.Println(4 >= 6)

	// Reflect
	fmt.Println(reflect.TypeOf(43))
	fmt.Println(reflect.TypeOf(3.1415))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("Hello Go!"))
}
