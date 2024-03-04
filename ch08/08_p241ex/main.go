package main

import "fmt"

type student struct {
  name string
  grade float64
}

func printInfo(s student) {
  fmt.Println("Name:", s.name)
  fmt.Println("Grade:", s.grade)
}

func main() {
  var s student
  s.name = "Audience kayela"
  s.grade = 92.3
  printInfo(s)
}
