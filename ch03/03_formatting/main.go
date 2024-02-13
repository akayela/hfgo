package main

import (
	"fmt"
)

func main() {
	var width, height, area float64
	width = 4.2
	height = 3.0
	area = width * height
	fmt.Printf("liters needed: %0.2f \n", area/10.0)
}
