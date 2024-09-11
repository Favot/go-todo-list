package task

import (
	"time"
)

type Task struct {
	ID          int
	Description string
	Completed   bool
	CreatedAt   time.Time
}

var (
	taskList []*Task
	nextID   int
)

func AddTask(description string) *Task {
	nextID++
	task := &Task{
		ID:          nextID,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}
	taskList = append(taskList, task)
	return task
}

func GetAllTask() []*Task {
	return taskList
}

func GetTaskByID(id int) *Task {
	for _, task := range taskList {
		if task.ID == id {
			return task
		}
	}
	return nil
}

func (t *Task) Complete() {
	t.Completed = true
}
