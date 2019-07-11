package cmd

import (
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	rootCmd.AddCommand(rmCmd)
}

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a task from the task list by index. Use `tkm stat` for indexes.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			Bail(ErrOnlyOneRm)
		}

		i, err := strconv.Atoi(args[0])

		if err != nil {
			Bail(ErrNAN)
		}

		ts, err := LoadTaskSet()

		must(err)
		must(ts.RemoveTask(i))
		must(UpdateTaskSet(ts))

		PrintTaskSet(ts)
	},
}
