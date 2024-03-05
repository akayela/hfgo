package main

import "fmt"

type myStruct struct {
  myField int
}

func main() {
  var value myStruct
  value.myField = 3
  var fieldPointer *myStruct = &value
  fmt.Println(fieldPointer.myField)
}
