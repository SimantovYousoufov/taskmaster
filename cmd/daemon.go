package cmd

import (
	"fmt"
	"github.com/0xAX/notificator"
	"github.com/simantovyousoufov/taskmaster/data"
	"github.com/spf13/cobra"
	"log"
	"time"
)

func init() {
	rootCmd.AddCommand(daemonCmd)
}

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Run the taskmaster daemon",
	Run: func(cmd *cobra.Command, args []string) {
		ticker := time.NewTicker(time.Minute)

		for range ticker.C {
			err := FindOutstandingTasks()

			if err != nil {
				log.Printf("Error: %s", err)
			}
		}
	},
}

func FindOutstandingTasks() error {
	ts, err := LoadTaskSet()

	if err != nil {
		return err
	}

	if ! ts.MITTasks.IsEmpty() {
		IterateTaskList(ts.MITTasks)
	}

	if ! ts.TodoTasks.IsEmpty() {
		IterateTaskList(ts.TodoTasks)
	}

	return nil
}

func IterateTaskList(tl data.TaskList) {
	for _, t := range tl.Tasks {
		if ! t.IsOutstanding() {
			continue
		}

		notify := notificator.New(notificator.Options{
			AppName:     "TaskMaster",
		})

		err := notify.Push("Outstanding Task", fmt.Sprintf("Task Is Outstanding: %s", t.Content), "", notificator.UR_NORMAL)

		if err != nil {
			log.Printf("Error: %s", err)
		}
	}
}
