package main 

import "fmt"

type Liters float64
type Gallons float64

func ToLiters(g Gallons) Liters {
  return Liters(g * 3.785)
}

func ToGallons(l Liters) Gallons {
  return Gallons(l * 0.264)
}

func main() {
  carFuel := Gallons(1.2)
  busFuel := Liters(2.5)
  carFuel += ToGallons(Liters(8.0))
  busFuel += ToLiters(Gallons(30.0))
  fmt.Println(carFuel)
}
