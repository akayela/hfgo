package main

import (
    "log"
    "net/http"
)

func write(w http.ResponseWriter, message string) {
    _, err := w.Write([]byte(message))
    if err != nil {
        log.Fatal(err)
    }
}

func englishHandler(w http.ResponseWriter, r *http.Request) {
    write(w, "Hello, Web")
}

func frenchHandler(w http.ResponseWriter, r *http.Request) {
    write(w, "Salut, Web")
}


func hindiHandler(w http.ResponseWriter, r *http.Request) {
    write(w, "Namaste, Web")
}

func main() {
    http.HandleFunc("/english", englishHandler)
    http.HandleFunc("/french", frenchHandler)
    http.HandleFunc("/hindi", hindiHandler)
    err := http.ListenAndServe("0.0.0.0:8001", nil)
    log.Fatal(err)
}
