package main

import (
	"fmt"
)

func paintNeeded(width, height float64) (float64, error) {
	if width <= 0 {
		return 0, fmt.Errorf("width cannot be a negative value")
	} else if height <= 0 {
		return 0, fmt.Errorf("height cannot be a negative value")
	} else {
		area := width * height
		return area / 10.0, nil
	}
}

func main() {
	var amount, total float64
	amount, _ = paintNeeded(4.2, 3.0)
	fmt.Printf("%0.2f liters needed \n", amount)
	total += amount

}
