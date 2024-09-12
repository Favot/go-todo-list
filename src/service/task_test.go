package service

import (
	"go-todo-list/src/models"
	"testing"
	"time"
)

type MockTaskRepository struct {
	tasks []*models.Task
}

func (m *MockTaskRepository) AddTask(task *models.Task) error {
	m.tasks = append(m.tasks, task)
	return nil
}

func (m *MockTaskRepository) GetAllTasks() ([]*models.Task, error) {
	return m.tasks, nil
}

func (m *MockTaskRepository) GetTaskByID(id int) (*models.Task, error) {
	for _, task := range m.tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return nil, nil
}

func (m *MockTaskRepository) SaveTasks(tasks []*models.Task) error {
	m.tasks = tasks
	return nil
}

func (m *MockTaskRepository) DeleteTask(id int) error {
	for i, task := range m.tasks {
		if task.ID == id {
			m.tasks = append(m.tasks[:i], m.tasks[i+1:]...)
			return nil
		}
	}
	return nil
}

func (m *MockTaskRepository) SaveTask(task *models.Task) error {
	for i, t := range m.tasks {
		if t.ID == task.ID {
			m.tasks[i] = task
			return nil
		}
	}
	return nil
}

func TestGetTaskByID(t *testing.T) {
	mockRepo := &MockTaskRepository{
		tasks: []*models.Task{
			{ID: 1, Description: "Task 1", Completed: false, CreatedAt: time.Now()},
			{ID: 2, Description: "Task 2", Completed: false, CreatedAt: time.Now()},
		},
	}
	service := NewTaskService(mockRepo)

	task, err := service.GetTaskByID(1)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if task == nil {
		t.Fatalf("expected task, got nil")
	}
	if task.ID != 1 {
		t.Errorf("expected task ID to be 1, got %d", task.ID)
	}
	if task.Description != "Task 1" {
		t.Errorf("expected task description to be 'Task 1', got %s", task.Description)
	}
}

func TestAddTask(t *testing.T) {
	mockRepo := &MockTaskRepository{
		tasks: []*models.Task{},
	}
	service := NewTaskService(mockRepo)

	task, err := service.AddTask("New Task")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if task == nil {
		t.Fatalf("expected task, got nil")
	}
	if task.ID != 1 {
		t.Errorf("expected task ID to be 1, got %d", task.ID)
	}
	if task.Description != "New Task" {
		t.Errorf("expected task description to be 'New Task', got %s", task.Description)
	}
	if task.Completed {
		t.Errorf("expected task to be incomplete, got completed")
	}
}

func TestCompleteTask(t *testing.T) {
	mockRepo := &MockTaskRepository{
		tasks: []*models.Task{
			{ID: 1, Description: "Task 1", Completed: false, CreatedAt: time.Now()},
		},
	}
	service := NewTaskService(mockRepo)

	err := service.CompleteTask(1)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	task, err := service.GetTaskByID(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if task == nil {
		t.Fatalf("expected task, got nil")
	}
	if !task.Completed {
		t.Errorf("expected task to be completed, got incomplete")
	}
}

func TestDeleteTask(t *testing.T) {
	mockRepo := &MockTaskRepository{
		tasks: []*models.Task{
			{ID: 1, Description: "Task 1", Completed: false, CreatedAt: time.Now()},
		},
	}
	service := NewTaskService(mockRepo)

	err := service.DeleteTask(1)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	task, err := service.GetTaskByID(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if task != nil {
		t.Errorf("expected task to be deleted, got %v", task)
	}
}
