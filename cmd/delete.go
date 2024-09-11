package cmd

import (
	"fmt"
	"go-todo-list/src/repository"
	"go-todo-list/src/service"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long:  "Delete a task by its ID",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide the ID of the task to delete")
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid ID:", args[0])
			return
		}
		repo := repository.NewCSVTaskRepository()
		taskService := service.NewTaskService(repo)
		deleteTask(taskService, id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deleteTask(service *service.TaskService, id int) {
	err := service.DeleteTask(id)
	if err != nil {
		fmt.Println("Error deleting task:", err)
		return

	}
	fmt.Println("Task deleted:", id)
}
