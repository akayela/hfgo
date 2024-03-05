package main

import "fmt"

type subscriber struct {
  name string
  rate float64
  active bool
}

func printInfo(s *subscriber) {
  fmt.Println("Name:", s.name)
  fmt.Println("Monthly rate:", s.rate)
  fmt.Println("Active", s.active)
}

func defaultSubscriber(name string) *subscriber {
  var s subscriber
  s.name = name
  s.rate = 5.99
  s.active = true
  return &s
}

func main() {
  pointerToSubscriber1 := defaultSubscriber("Audience Kayela")
  pointerToSubscriber1.rate = 4.99
  printInfo(pointerToSubscriber1)

  pointerToSubscriber2 := defaultSubscriber("Joana Kayela")
  printInfo(pointerToSubscriber2)
}
