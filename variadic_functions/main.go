package main

import "fmt"

func severalInts(numbers ...int) {
	fmt.Println(numbers)
}

func severalStrings(strings ...string) {
	fmt.Println(strings)
}

func mix(num int, flag bool, strings ...string) {
	fmt.Println(num, flag, strings)
}

func main() {
	severalInts(1)
	severalInts(1, 2, 4, 6)
	severalStrings("e")
	severalStrings("a", "b", "c", "d", "e")
	mix(1, true, "a", "b")
	mix(2, false, "d", "e")
}
