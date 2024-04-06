package main

import (
    "fmt"
    "time"
)

func a() {
    for i := 0; i < 50; i++ {
        fmt.Print("a", i)
    }
}

func b() {
    for i := 0; i < 50; i++ {
        fmt.Print("b", i)
    }
}

func main() {
    go a()
    go b()
    time.Sleep(time.Second)
    fmt.Println("end main()")
}
