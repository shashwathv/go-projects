package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/shashwathv/todo_app/internal/task"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Deletes a task",
	Args:  cobra.ExactArgs(1),
	Run:   runDelete,
}

func runDelete(cmd *cobra.Command, args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := task.DeleteTask(id); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("Task deleted")
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
