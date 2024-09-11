# Simple Todo Task Manager

This is a simple CLI-based Todo Task Manager built with Go. The project is designed to help you learn Go and practice using the Repository Pattern. The CLI part of the application is implemented using Cobra.

## Features

- Add a new task
- List all tasks
- Mark a task as complete

## Getting Started

### Prerequisites

- Go 1.16 or later

### Installation

1. Clone the repository:

```sh
git clone https://github.com/yourusername/your_project.git
cd your_project
```

2. Install dependencies:

```sh
go mod tidy
```

3. Run the application:

```sh
go run main.go
```

## Usage

### Add a new task

```sh
go run main.go add "Buy groceries"
```

### List all tasks

```sh
go run main.go list
```

### Mark a task as complete

```sh
go run main.go complete 1
```

### Delete a task

```sh
go run main.go delete 1
```
