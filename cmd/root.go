package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var verbose bool
var path string

var rootCmd = &cobra.Command{
	Use:     "arr",
	Aliases: []string{"r"},
	Short:   "R from the command line",
	Long:    `Call common R functions from your terminal`,
	Version: "0.0.1",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", true, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&path, "path", "p", "", "path to R installation")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
