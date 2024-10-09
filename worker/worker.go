package main

import (
    "fmt"
    "time"
)

func registerWorker() {
    fmt.Println("Worker registered.")
    // Send registration info to the server
}

func processTask(task Task) {
    fmt.Printf("Processing Task %d: %s\n", task.ID, task.Detail)
    time.Sleep(3 * time.Second) // Simulate work
    fmt.Printf("Task %d completed.\n", task.ID)
}
