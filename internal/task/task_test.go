package task

import (
	"testing"
)

func resetTaskList() {
	taskList = []*Task{}
	nextID = 0
}

func TestNewTaskItem(t *testing.T) {
	resetTaskList()
	task := AddTask("Test task")
	if task.ID != 1 || task.Description != "Test task" || task.Completed != false {
		t.Errorf("NewTaskItem() failed, got: %v", task)
	}
}

func TestComplete(t *testing.T) {
	resetTaskList()
	task := AddTask("Test task")
	task.Complete()
	if !task.Completed {
		t.Errorf("Complete() failed, task should be completed")
	}

	// Verify the task is updated in the slice
	updatedTask := GetTaskByID(task.ID)
	if updatedTask == nil || !updatedTask.Completed {
		t.Errorf("GetTaskByID() failed, task should be completed in the slice")
	}
}

func TestGetAllTask(t *testing.T) {
	resetTaskList()

	task1 := AddTask("Task 1")
	task2 := AddTask("Task 2")

	tasks := GetAllTask()
	if len(tasks) != 2 {
		t.Errorf("GetAllTask() failed, expected 2 tasks, got: %d", len(tasks))
	}

	if tasks[0] != task1 || tasks[1] != task2 {
		t.Errorf("GetAllTask() failed, tasks do not match")
	}
}
