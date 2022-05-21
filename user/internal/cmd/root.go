package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string //nolint:deadcode,varcheck,unused

// var cmdHandler = cli.NewHandler(env.Env)

// RootCmd This represents the base command when called without any subcommands.
var RootCmd = &cobra.Command{
	Use:   "worker",
	Short: "Worker CLI",
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// TODO: consider move configs to here
}
