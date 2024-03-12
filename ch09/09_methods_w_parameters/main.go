package main

import "fmt"

type MyType string

func (m MyType) MethodWithParams(number int, flag bool) {
  fmt.Println(m)
  fmt.Println(number)
  fmt.Println(flag)
}

func (m MyType) WithReturn() int {
  return len(m)
}

func main() {
  value := MyType("MyType value")
  value.MethodWithParams(4,true)

  anotherValue := MyType("Another value of my type")
  fmt.Println(anotherValue.WithReturn())
}
