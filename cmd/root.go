package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var verbose bool
var path string
var arg1 string
var arg2 string
var arg3 string
var arg1s string
var arg2s string
var arg3s string

var rootCmd = &cobra.Command{
	Use: "arr",
	ValidArgs: []string{
		"devtools",
		"usethis",
		"script",
		"run",
		"packer",
		"rsconnect",
		"completion",
	},
	Short: "R from the command line",
	Long: `Call common R functions from your terminal

Examples:

	# arr package function
	arr devtools document check install
	arr usethis use_mit_license use_pipe
	arr packer bundle
	arr rsConnect writeManifest

	Passing arguments.

	# bare arguments
	arr usethis use_gpl_license -1=3 -2=TRUE

	# as quote string
	arr packer bundle --arg1s="prod"
	`,
	Version: "0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&path, "path", "p", "", "path to R installation")
	rootCmd.PersistentFlags().StringVarP(&arg1, "arg1", "1", "", "first argument to pass to the function call")
	rootCmd.PersistentFlags().StringVarP(&arg2, "arg2", "2", "", "second argument to pass to the function call")
	rootCmd.PersistentFlags().StringVarP(&arg3, "arg3", "3", "", "third argument to pass to the function call")

	rootCmd.PersistentFlags().StringVarP(&arg1s, "arg1s", "", "", "first argument as quoted string to pass to the function call")
	rootCmd.PersistentFlags().StringVarP(&arg2s, "arg2s", "", "", "second argument as quoted string to pass to the function call")
	rootCmd.PersistentFlags().StringVarP(&arg3s, "arg3s", "", "", "third argument as quoted string to pass to the function call")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
