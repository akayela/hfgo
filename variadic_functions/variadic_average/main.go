package main

import (
	"fmt"
	"log"
)

func average(numbers ...float64) (float64, error) {
	var sum float64 = 0
	for i, number := range numbers {
		if numbers[i] < 0 {
			return 0, fmt.Errorf("invalid value: %0.2f", numbers[i])
		}
		sum += number
	}
	return sum / float64(len(numbers)), nil
}

func main() {
	avg1, err := average(100, 50)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(avg1)
	avg2, err := average(90.7, 98.2, 98.5, -92.3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(avg2)
}
