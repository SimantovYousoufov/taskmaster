package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(statCmd)
}

var statCmd = &cobra.Command{
	Use:   "stat",
	Short: "View the task set. Aliases: ls, list, all",
	Aliases: []string{"ls", "list", "all"},
	Run: func(cmd *cobra.Command, args []string) {
		ts, err := LoadTaskSet()
		must(err)

		PrintTaskSet(ts)
	},
}
