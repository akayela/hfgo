package main

import "fmt"

func main() {
	// initialize a map
	ranks := make(map[string]int)
	ranks["Bronze"] = 3
	ranks["Silver"] = 2
	ranks["Gold"] = 1
	fmt.Println(ranks["Bronze"])
	fmt.Println(ranks["Gold"])

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["Li"] = "Lithium"
	fmt.Println(elements["H"])
	fmt.Println(elements["Li"])

	isPrime := make(map[int]bool)
	isPrime[4] = false
	isPrime[7] = true
	fmt.Println(isPrime[4])
	fmt.Println(isPrime[7])

	teamKayela := map[string]string{
		"dad":      "Audience",
		"mom":      "Joana",
		"daughter": "Nicole",
		"son":      "Donne",
	}
	fmt.Println(teamKayela["dad"])

	// initialize a map and a bool to check assignment
	counters := map[string]int{"a": 3, "b": 0}
	var value int
	var ok bool
	value, ok = counters["a"]
	fmt.Println(value, ok)
	value, ok = counters["b"]
	fmt.Println(value, ok)
	value, ok = counters["c"]
	fmt.Println(value, ok)

	teamCorner := map[string]string{
		"lead":  "Audience",
		"tech1": "Thabo",
		"tech2": "Abel",
		"tech3": "Tebogo",
	}

	tech1, ok := teamCorner["tech1"]
	tech2, ok := teamCorner["tech2"]
	fmt.Println(tech1, ok)
	fmt.Println(tech2, ok)
	delete(teamCorner, "tech1")
	tech1, ok = teamCorner["tech1"]
	fmt.Println(tech1, ok)

	status("Alma")
	status("Carl")
}
func status(name string) {
	grades := map[string]float64{"Alma": 0, "Rohit": 86.5}
	grade, ok := grades[name]
	if !ok {
		fmt.Printf("No grade assigned for %s\n", name)
	} else if grade < 60 {
		fmt.Printf("%s is failing\n", name)
	}
}
