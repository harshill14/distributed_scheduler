package main

import "fmt"

type Worker struct {
    ID   int
    Busy bool
}

var workers []Worker

func assignTaskToWorker(task Task) {
    for i, worker := range workers {
        if !worker.Busy {
            fmt.Printf("Assigning Task %d to Worker %d\n", task.ID, worker.ID)
            workers[i].Busy = true
            // Simulate task processing
            task.Status = "In Progress"
            return
        }
    }
    fmt.Println("No available workers. Task queued.")
}
