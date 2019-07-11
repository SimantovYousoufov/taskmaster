package cmd

import (
	"fmt"
	"github.com/simantovyousoufov/taskmaster/data"
	"github.com/spf13/cobra"
	"strings"
)

func init() {
	rootCmd.AddCommand(todoCmd)
}

var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "Add a task to the TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		ts, err := LoadTaskSet()
		must(err)

		newTodo := strings.Join(args, " ")
		err = ts.AddTask(data.TodoTask, newTodo)

		if err == data.ErrAtTaskLimit {
			PrintTaskSet(ts)
			Bail(err)
		}

		must(err)
		must(UpdateTaskSet(ts))

		fmt.Printf("Successfully added new TODO. Current tasks:\n")
		PrintTaskSet(ts)
	},
}
