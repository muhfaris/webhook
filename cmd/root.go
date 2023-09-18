package cmd

import (
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use: "webhook",
}

var cfgFile string

func Run() {
	restCommand.PersistentFlags().StringVar(&cfgFile, "config", "config.json", "config file (default is config.json)")
	rootCommand.AddCommand(restCommand)
	if err := rootCommand.Execute(); err != nil {
		panic(err)
	}
}
