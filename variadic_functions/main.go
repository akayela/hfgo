package main

import (
	"fmt"
	"math"
)

func severalInts(numbers ...int) {
	fmt.Println(numbers)
}

func severalStrings(strings ...string) {
	fmt.Println(strings)
}

func mix(num int, flag bool, strings ...string) {
	fmt.Println(num, flag, strings)
}

func maximum(numbers ...float64) float64 {
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func main() {
	severalInts(1)
	severalInts(1, 2, 4, 6)
	severalStrings("e")
	severalStrings("a", "b", "c", "d", "e")
	mix(1, true, "a", "b")
	mix(2, false, "d", "e")
	fmt.Println(maximum(71.8, 56.2, 89.3))
	fmt.Println(maximum(90.7, 89.7, 98.3, 92.3))
}
