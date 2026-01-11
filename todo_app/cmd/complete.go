package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/shashwathv/todo_app/internal/task"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete <id>",
	Short: "Mark a task as complete",
	Args:  cobra.ExactArgs(1),
	Run:   runComplete,
}

func runComplete(cmd *cobra.Command, args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := task.CompleteTask(id); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("Task marked as complete.")

}

func init() {
	rootCmd.AddCommand(completeCmd)
}
