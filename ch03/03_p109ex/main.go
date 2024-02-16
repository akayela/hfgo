package main

import "fmt"

func negate(myBoolean *bool) bool {
	*myBoolean = !*myBoolean
	return *myBoolean
}

func main() {
	truth := true
	negate(&truth)
	fmt.Println(truth)
	lies := false
	negate(&lies)
	fmt.Println(lies)
}
