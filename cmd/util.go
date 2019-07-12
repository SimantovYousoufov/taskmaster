package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/simantovyousoufov/taskmaster/data"
	"github.com/spf13/viper"
	"os"
)

func LoadTaskSet() (*data.TaskSet, error) {
	if ! viper.InConfig(AppDataKey) {
		return data.NewTaskSet(), nil
	}

	ts := &data.TaskSet{}
	storedJson := viper.GetString(AppDataKey)

	err := json.Unmarshal([]byte(storedJson), ts)

	if err != nil {
		return nil, errors.New("unable to read in task set")
	}

	ts.MITTasks.TaskLimit = viper.GetInt(MITLimitKey)
	ts.TodoTasks.TaskLimit = viper.GetInt(TodoLimitKey)

	return ts, err
}

func UpdateTaskSet(ts *data.TaskSet) error {
	tsData, _ := ts.ToJSON()

	viper.Set(AppDataKey, string(tsData))

	return viper.WriteConfig()
}

func PrintTaskSet(ts *data.TaskSet) {
	i := 0

	if ! ts.MITTasks.IsEmpty() {
		fmt.Printf("MITs:\n")

		PrintTaskList(ts.MITTasks, &i)
	}

	if ! ts.TodoTasks.IsEmpty() {
		fmt.Printf("TODOs:\n")

		PrintTaskList(ts.TodoTasks, &i)
	}
}

func PrintTaskList(tl data.TaskList, i *int) {
	for _, t := range tl.Tasks {
		fmt.Printf("%d) %s", *i, t.Content)

		fmt.Printf("\n")
		*i++
	}

	fmt.Printf("\n")
}

func PrintError(err error) {
	color.Red(fmt.Sprintf("Error: %s\n\n", err))
}

func Bail(err error) {
	fmt.Printf("Error: %s\n", err)
	os.Exit(1)
}

func must(err error) {
	if err != nil {
		Bail(err)
	}
}
