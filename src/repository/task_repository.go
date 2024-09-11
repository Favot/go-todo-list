package repository

import (
	"go-todo-list/src/models"
)

type TaskRepository interface {
	AddTask(task *models.Task) error
	GetAllTasks() ([]*models.Task, error)
	GetTaskByID(id int) (*models.Task, error)
	SaveTasks(tasks []*models.Task) error
}
