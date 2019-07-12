package cmd

import (
	"github.com/simantovyousoufov/taskmaster/data"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	rootCmd.AddCommand(mvCmd)
}

var mvCmd = &cobra.Command{
	Use:   "mv",
	Short: "Move a task from the MIT list to the TODOs list.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			Bail(ErrOnlyOneMv)
		}

		i, err := strconv.Atoi(args[0])

		if err != nil {
			Bail(ErrNAN)
		}

		ts, err := LoadTaskSet()

		if i > data.MITLimit || i > len(ts.MITTasks.Tasks) {
			Bail(ErrInvalidMv)
		}

		t := ts.MITTasks.Tasks[i]

		must(err)
		must(ts.RemoveTask(i))
		must(ts.AddTask(data.TodoTask, t.Content))
		must(UpdateTaskSet(ts))

		PrintTaskSet(ts)
	},
}
