package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "bytes"
)

func submitTask(task string) {
    url := "http://localhost:8080/submit"
    taskRequest := map[string]string{"task": task}
    jsonValue, _ := json.Marshal(taskRequest)
    _, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
    if err != nil {
        fmt.Println("Error submitting task:", err)
    } else {
        fmt.Println("Task submitted successfully")
    }
}

func main() {
    task := "Analyze data"
    submitTask(task)
}
