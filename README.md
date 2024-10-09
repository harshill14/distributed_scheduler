# Distributed Task Scheduler

A Distributed Task Scheduler built with Go, designed to enable asynchronous task distribution and processing across multiple worker nodes. This system allows clients to submit tasks to a server, which then assigns them to available worker nodes for processing, with results sent back to the client.

## Features

- **Task Server**: Accepts tasks from clients, queues them, and assigns them to available worker nodes.
- **Worker Nodes**: Processes tasks asynchronously and reports results back to the server.
- **Client**: Submits tasks to the server and retrieves results upon completion.
- **Scalability**: Supports the addition of multiple worker nodes for distributed processing.
- **Asynchronous Processing**: Ensures efficient task handling by leveraging worker nodes that operate independently.

## Directory Structure

```bash
distributed_scheduler/
├── client/
│   ├── client.go
│   └── task.go
├── server/
│   ├── server.go
│   ├── task_manager.go
│   └── worker_manager.go
└── worker/
    ├── worker.go
    └── task_processor.go
