package main

import "fmt"

type Liters float64
type Milliliters float64
type Gallons float64

func (l Liters) ToGallons() Gallons {
    return Gallons(l * 0.264)
}

func (m Milliliters) ToGallons() Gallons {
    return Gallons(m * 0.000264)
}

func (g Gallons) ToLiters() Liters {
    return Liters(g / 0.264)
}

func (m Milliliters) ToLiters() Liters {
    return Liters(m / 1000)
}
