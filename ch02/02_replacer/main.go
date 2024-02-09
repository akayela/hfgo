package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# R#cks!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}
