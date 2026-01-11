package cmd

import (
	"fmt"
	"os"

	"github.com/shashwathv/todo_app/internal/task"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <descriptor>",
	Short: "Add a new task",
	Args:  cobra.ExactArgs(1),
	Run:   runAdd,
}

func runAdd(cmd *cobra.Command, args []string) {
	if err := task.AddTask(args[0]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("Task added.")
}

func init() {
	rootCmd.AddCommand(addCmd)
}
