package cmd

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strconv"
)

func init() {
	rootCmd.AddCommand(configCmd)
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure Taskmaster",
	Run: func(cmd *cobra.Command, args []string) {
		mitLimit := promptForInt("MIT Task Limit (number)")
		todoLimit := promptForInt("Todo Task Limit (number)")

		viper.Set(MITLimitKey, mitLimit)
		viper.Set(TodoLimitKey, todoLimit)

		fmt.Println("Configuration was successfully updated")

		viper.WriteConfig()
	},
}

func promptForInt(prompt string) int64 {
	intPrompt := promptui.Prompt{
		Label:   prompt,
		Validate: validateIntInput,
	}

	data, err := intPrompt.Run()
	must(err)

	parsed, err := strconv.ParseInt(data, 10, 64)
	must(err)

	return parsed
}

func validateIntInput(input string) error {
	_, err := strconv.ParseInt(input, 10, 64)

	if err != nil {
		return ErrNAN
	}

	return nil
}
