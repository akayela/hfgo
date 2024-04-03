package main

import "fmt"

func find(item string, slice []string) bool {
    for _, sliceItem := range slice {
        if item == sliceItem {
            return true
        }
    }
    return false
}

type Refrigerator []string

func (r Refrigerator) Open() {
    fmt.Println("Opening refrigerator")
}

func (r Refrigerator) Close() {
    fmt.Println("Closing refrigerator")
}

func (r Refrigerator) FindFood(food string) error {
    r.Open()
    if find(food, r) {
        fmt.Println("Found", food)
    }else{
        return fmt.Errorf("%s not found", food)
    }
    r.Close()
    return nil
}
