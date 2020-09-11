package cmd

import (
	"fmt"
	"os"

	"github.com/cli_task/task/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("something went wrong:", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("you have no tasks to complete! why not take a vacation")
			return
		}
		fmt.Println("you have the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s, Key = %d\n", i+1, task.Value, task.Key)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
