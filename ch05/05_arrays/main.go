package main

import "fmt"

func main() {
	var notes [7]string
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	notes[3] = "fa"
	notes[4] = "so"
	notes[5] = "la"
	notes[6] = "ti"
	for _, note := range notes {
		fmt.Println(note)
	}
}
