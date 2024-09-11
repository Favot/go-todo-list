package service

import (
	"go-todo-list/src/models"
	"go-todo-list/src/repository"
	"time"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) AddTask(description string) (*models.Task, error) {
	nextID := s.getNextID()
	task := &models.Task{
		ID:          nextID,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}
	if err := s.repo.AddTask(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) GetAllTasks() ([]*models.Task, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) GetTaskByID(id int) (*models.Task, error) {
	return s.repo.GetTaskByID(id)
}

func (s *TaskService) CompleteTask(id int) error {
	task, err := s.repo.GetTaskByID(id)
	if err != nil {
		return err
	}
	if task == nil {
		return nil
	}
	task.Completed = true
	tasks, err := s.repo.GetAllTasks()
	if err != nil {
		return err
	}
	return s.repo.SaveTasks(tasks)
}

func (s *TaskService) getNextID() int {
	tasks, err := s.repo.GetAllTasks()
	if err != nil {
		panic(err)
	}
	return len(tasks) + 1
}
