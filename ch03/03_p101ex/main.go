package main

import (
	"errors"
	"fmt"
)

func devide(dividend, divisor float64) (float64, error) {
	if divisor == 0.0 {
		return 0.0, errors.New("can't devide by zero")
	}
	return dividend / divisor, nil
}

func main() {
	quotient, err := devide(5.6, 0.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.2f\n", quotient)
	}
}
