package cmd

import (
	"fmt"
	"go-todo-list/src/repository"
	"go-todo-list/src/service"
	"strconv"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete a task",
	Long:  "Complete a task by its ID",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide the ID of the task to complete")
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid ID:", args[0])
			return
		}
		repo := repository.NewCSVTaskRepository()
		taskService := service.NewTaskService(repo)
		completeTask(taskService, id)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}

func completeTask(service *service.TaskService, id int) {
	err := service.CompleteTask(id)
	if err != nil {
		fmt.Println("Error completing task:", err)
		return
	}
	fmt.Println("Task completed:", id)
}
