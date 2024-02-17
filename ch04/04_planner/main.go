package main

import (
	"dates"
	"fmt"
)

func main() {
	days := 3
	fmt.Println("Your appointement is in", days, "days")
	fmt.Println("With a follow up in", days+dates.DaysInWeek, "days")
}
