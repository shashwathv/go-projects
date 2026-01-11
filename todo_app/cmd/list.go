package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/mergestat/timediff"
	"github.com/shashwathv/todo_app/internal/task"
	"github.com/spf13/cobra"
)

var showAll bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Args:  cobra.ExactArgs(0),
	Run:   runList,
}

func runList(cmd *cobra.Command, args []string) {
	tasks, err := task.ListTasks()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	w := tabwriter.NewWriter(
		os.Stdout,
		0, 0, 2, ' ',
		0,
	)

	// Print header
	if showAll {
		fmt.Fprintln(w, "ID\tTask\tCreated\tDone")
	} else {
		fmt.Fprintln(w, "ID\tTask\tCreated")
	}

	// Print rows
	for _, t := range tasks {
		if !showAll && t.IsComplete {
			continue
		}

		if showAll {
			fmt.Fprintf(
				w,
				"%d\t%s\t%s\t%v\n",
				t.ID,
				t.Description,
				timediff.TimeDiff(t.CreatedAt),
				t.IsComplete,
			)
		} else {
			fmt.Fprintf(
				w,
				"%d\t%s\t%s\n",
				t.ID,
				t.Description,
				timediff.TimeDiff(t.CreatedAt),
			)
		}
	}

	w.Flush()
}

func init() {
	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Show all tasks")
	rootCmd.AddCommand(listCmd)
}
