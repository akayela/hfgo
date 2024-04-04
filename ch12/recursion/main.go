package main 

import "fmt"

func count(start int, end int) {
    if start < end {
        fmt.Println("Starting at: ", start)
        count(start+1, end)
    }
}

func main() {
    count(1, 10)
}
