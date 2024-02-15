package main

import (
	"fmt"
	"log"
)

func paintNeeded(width, height float64) (float64, error) {
	if width <= 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	} else if height <= 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	} else {
		area := width * height
		return area / 10.0, nil
	}
}

func main() {
	var total float64
	amount, err := paintNeeded(4.2, 3.0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%0.2f liters needed \n", amount)
	total += amount

}
