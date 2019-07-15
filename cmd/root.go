package cmd

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/simantovyousoufov/taskmaster/data"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.taskmaster.json)")

	viper.SetDefault(MITLimitKey, data.MITLimit)
	viper.SetDefault(TodoLimitKey, data.TodoLimit)
}

func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigFile(path.Join(home, ProjectConfigFile))
	}

	if _, err := os.Stat(path.Join(home, ProjectConfigFile)); os.IsNotExist(err) {
		f, err := os.Create(path.Join(home, ProjectConfigFile))
		must(err)

		f.Close()

		viper.WriteConfig()
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "tkm",
	Short: "Taskmaster is a better way to manage your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		statCmd.Run(cmd, args)
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("Config file changed:", e.Name)
		})
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		viper.WriteConfig()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
