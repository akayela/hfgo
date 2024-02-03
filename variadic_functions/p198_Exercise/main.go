package main

import (
	"fmt"
	"log"
)

func sum(numbers ...int) (int, error) {
	var sum int = 0
	for i, number := range numbers {
		if numbers[i] < 0 {
			return 0, fmt.Errorf("invalid input: %d", numbers[i])
		}
		sum += number
	}
	return sum, nil
}

func main() {
	sum1, err := (sum(9, 7))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum1)
	sum2, err := (sum(1, 4, 2))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum2)
}
