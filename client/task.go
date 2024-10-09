package client

import (
    "encoding/json"
    "fmt"
    "time"
)

// Task represents a client task to be submitted to the server
type Task struct {
    ID          int       `json:"id"`
    Detail      string    `json:"detail"`
    Status      string    `json:"status"`
    SubmittedAt time.Time `json:"submitted_at"`
}

// NewTask creates a new task with a given detail
func NewTask(detail string) *Task {
    return &Task{
        ID:          generateTaskID(),
        Detail:      detail,
        Status:      "Pending",
        SubmittedAt: time.Now(),
    }
}

// generateTaskID generates a unique task ID for each task
// In a distributed system, you might use a UUID or a centralized ID service
var taskIDCounter int

func generateTaskID() int {
    taskIDCounter++
    return taskIDCounter
}

// UpdateStatus updates the task status
func (t *Task) UpdateStatus(status string) {
    t.Status = status
    fmt.Printf("Task %d status updated to %s\n", t.ID, t.Status)
}

func (t *Task) ToJSON() ([]byte, error) {
    jsonData, err := json.Marshal(t)
    if err != nil {
        return nil, fmt.Errorf("failed to serialize task %d: %v", t.ID, err)
    }
    return jsonData, nil
}


func (t *Task) PrintSummary() {
    fmt.Printf("Task Summary:\n")
    fmt.Printf("ID: %d\n", t.ID)
    fmt.Printf("Detail: %s\n", t.Detail)
    fmt.Printf("Status: %s\n", t.Status)
    fmt.Printf("Submitted At: %s\n", t.SubmittedAt.Format(time.RFC1123))
}

// Example
func main() {
    task := NewTask("Analyze data set for trends")
    task.PrintSummary()

    taskJSON, err := task.ToJSON()
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Task JSON:", string(taskJSON))
    }

    task.UpdateStatus("In Progress")
    task.PrintSummary()
}
