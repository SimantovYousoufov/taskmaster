package cmd

import (
	"fmt"
	"github.com/simantovyousoufov/taskmaster/data"
	"github.com/spf13/cobra"
	"strings"
)

func init() {
	rootCmd.AddCommand(mitCmd)
}

var mitCmd = &cobra.Command{
	Use:   "mit",
	Short: "Add a task to the MIT list",
	Run: func(cmd *cobra.Command, args []string) {
		ts, err := LoadTaskSet()
		must(err)

		newMit := strings.Join(args, " ")
		err = ts.AddTask(data.MITTask, newMit)

		if err == data.ErrAtTaskLimit {
			PrintTaskSet(ts)
			Bail(err)
		}

		must(err)
		must(UpdateTaskSet(ts))

		fmt.Printf("Successfully added new MIT. Current tasks:\n")
		PrintTaskSet(ts)
	},
}
