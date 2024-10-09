package main

import (
    "fmt"
)

type Task struct {
    ID     int
    Detail string
    Status string
}

var taskQueue []Task

func addTask(detail string) Task {
    task := Task{ID: len(taskQueue) + 1, Detail: detail, Status: "Pending"}
    taskQueue = append(taskQueue, task)
    return task
}
