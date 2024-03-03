package main

import "fmt"

type part struct {
	description string
	count       int
	location    string
}

func showInfo(p part) {
	fmt.Println("Decription: ", p.description)
	fmt.Println("Count: ", p.count)
}

func minimumOrder(description string, count int) part {
	var p part
	p.description = description
	p.count = count
	p.location = "Main warehouse"
	return p
}

func main() {
	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 32
	showInfo(bolts)

	p := minimumOrder("Filter wrench", 20)
	fmt.Println(p.description, p.count, p.location)

}
