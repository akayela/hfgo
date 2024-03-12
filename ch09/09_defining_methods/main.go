package main

import "fmt"

type MyType string 

func (m MyType) sayHi() {
  fmt.Println("Hi")
}

func main() {
  value := MyType("a MyType value")
  value.sayHi()
  anotherValue := MyType("another MyType value")
  anotherValue.sayHi()
}
