package cmd

import (
	"fmt"
	"go-todo-list/internal/task"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `List all tasks in your todo list.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks := task.GetAllTask()
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
	},
}
