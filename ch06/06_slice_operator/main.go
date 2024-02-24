package main

import "fmt"

func main() {

	underLyingArray := [5]string{"a", "b", "c", "d", "e"}
	slice1 := underLyingArray[0:3]
	slice3 := underLyingArray[0:]
	fmt.Println(slice1)
	fmt.Println(underLyingArray[2:])
	underLyingArray[2] = "v"
	fmt.Println("After changing the underlying array:")
	fmt.Println(slice1)
	slice3 = append(slice3, "f")
	fmt.Println(slice3)

	i, j := 1, 4
	slice2 := underLyingArray[i:j]
	fmt.Println(slice2)
}
