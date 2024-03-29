package main

import (
	"fmt"
	"log"

	"github.com/akayela/keyboard"
)

func main() {
	fmt.Print("Enter a temperature in Farenheit: ")
	farenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celcius := (farenheit - 32) * 5 / 9
	fmt.Printf("%.2f degrees Celcius\n", celcius)
}
