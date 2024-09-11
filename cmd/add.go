package cmd

import (
	"fmt"
	"go-todo-list/src/repository"
	"go-todo-list/src/service"

	"github.com/spf13/cobra"
)

var description string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  `Add a new task to your todo list.`,
	Run: func(cmd *cobra.Command, args []string) {
		if description == "" {
			fmt.Println("Description is required to add a task")
			return
		}
		repo := repository.NewCSVTaskRepository()
		taskService := service.NewTaskService(repo)
		addTask(taskService, description)
	},
}

func init() {
	addCmd.Flags().StringVarP(&description, "description", "d", "", "Description of the task to add")
	rootCmd.AddCommand(addCmd)
}

func addTask(service *service.TaskService, description string) {
	task, err := service.AddTask(description)
	if err != nil {
		fmt.Println("Error adding task:", err)
		return
	}
	fmt.Println("Added task:", task)
}
