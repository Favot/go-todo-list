package cmd

import (
	"fmt"
	"go-todo-list/src/repository"
	"go-todo-list/src/service"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `List all tasks in your todo list.`,
	Run: func(cmd *cobra.Command, args []string) {
		repo := repository.NewCSVTaskRepository()
		taskService := service.NewTaskService(repo)
		listTasks(taskService)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listTasks(service *service.TaskService) {
	tasks, err := service.GetAllTasks()
	if err != nil {
		fmt.Println("Error getting tasks:", err)
		return
	}
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, t := range tasks {
		status := "Incomplete"
		if t.Completed {
			status = "Complete"
		}
		fmt.Printf("%d: %s [%s]\n", t.ID, t.Description, status)
	}
}
