package main

type Task struct {
    ID     int
    Detail string
    Status string
}

func handleTaskProcessing(task Task) {
    fmt.Println("Processing task:", task.Detail)
    // Custom processing logic
}
