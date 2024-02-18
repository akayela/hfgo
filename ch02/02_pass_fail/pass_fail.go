// pass_fail reports whether a grade is passing or failing.
package main

import (
	"fmt"
	"log"

	"github.com/akayela/keyboard"
)

func main() {
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if err != nil {
		log.Fatal(err)
	}
	if grade == 100 {
		status = "Perfect"
	} else if grade >= 60 {
		status = "Passing"
	} else {
		status = "Failing"
	}
	fmt.Printf("A grade 0f %.2f is %s\n", grade, status)
}
