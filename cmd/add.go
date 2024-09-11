package cmd

import (
	"fmt"

	"go-todo-list/internal/task"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  `Add a new task to your todo list.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a task description.")
			return
		}
		description := args[0]
		task := task.AddTask(description)
		fmt.Printf("Added task: %d - %s\n", task.ID, task.Description)
	},
}
