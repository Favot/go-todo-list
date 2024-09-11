package repository

import (
	"encoding/csv"
	"go-todo-list/src/models"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type CSVTaskRepository struct {
	csvFile string
}

func NewCSVTaskRepository() *CSVTaskRepository {
	return &CSVTaskRepository{
		csvFile: filepath.Join("database", "tasks.csv"),
	}
}

func (r *CSVTaskRepository) AddTask(task *models.Task) error {
	file, err := os.OpenFile(r.csvFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{
		strconv.Itoa(task.ID),
		task.Description,
		strconv.FormatBool(task.Completed),
		task.CreatedAt.Format(time.RFC3339),
	}
	return writer.Write(record)
}

func (r *CSVTaskRepository) GetAllTasks() ([]*models.Task, error) {
	file, err := os.Open(r.csvFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []*models.Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var tasks []*models.Task
	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		completed, _ := strconv.ParseBool(record[2])
		createdAt, _ := time.Parse(time.RFC3339, record[3])
		task := &models.Task{
			ID:          id,
			Description: record[1],
			Completed:   completed,
			CreatedAt:   createdAt,
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *CSVTaskRepository) GetTaskByID(id int) (*models.Task, error) {
	tasks, err := r.GetAllTasks()
	if err != nil {
		return nil, err
	}
	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return nil, nil
}

func (r *CSVTaskRepository) SaveTasks(tasks []*models.Task) error {
	file, err := os.Create(r.csvFile)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, task := range tasks {
		record := []string{
			strconv.Itoa(task.ID),
			task.Description,
			strconv.FormatBool(task.Completed),
			task.CreatedAt.Format(time.RFC3339),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	return nil
}

func (r *CSVTaskRepository) DeleteTask(id int) error {
	tasks, err := r.GetAllTasks()
	if err != nil {
		return err
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return r.SaveTasks(tasks)
		}
	}
	return nil
}
