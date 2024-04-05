package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func reportPanic() {
    p := recover()
    if p == nil {
        return
    } 
    err, ok := p.(error)
    if ok {
        fmt.Println(err)
    }
}

func scanDirectory (path string) {
    fmt.Println(path)
    files, err := os.ReadDir(path)
    if err != nil {
        panic (err)
    }
    for _, file := range files {
        filePath := filepath.Join(path, file.Name())
        if file.IsDir() {
            scanDirectory(filePath)
            } else {
                fmt.Println(filePath)
            }
        }
}

func main() {
    defer reportPanic()
    scanDirectory("go")
}
