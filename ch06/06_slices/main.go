package main

import (
	"fmt"
)

func main() {
	notes2 := []string{"do", "re", "mi", "fa", "so", "la", "ti"}
	var notes []string // declare a slice
	notes = make([]string, 7)
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	notes[3] = "fa"
	notes[4] = "so"
	notes[5] = "la"
	notes[6] = "ti"
	for _, note := range notes {
		fmt.Println(note, " ")
	}
	fmt.Println(notes[0], " ")

	for i := 6; i >= 0; i-- {
		fmt.Println(notes[i], " ")
	}

	fmt.Println(notes2[0], " ")
}
