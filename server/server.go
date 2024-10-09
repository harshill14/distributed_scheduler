package main

import (
    "net/http"
    "fmt"
)

func main() {
    http.HandleFunc("/submit", submitTaskHandler)
    fmt.Println("Server running on port 8080...")
    http.ListenAndServe(":8080", nil)
}
