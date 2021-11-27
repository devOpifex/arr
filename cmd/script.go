package cmd

import (
	"github.com/spf13/cobra"
)

var sourceCmd = &cobra.Command{
	Use:   "script",
	Short: "run scripts",
	Long:  `run R scripts`,
	Args:  cobra.MinimumNArgs(1),
	Run:   makeRun("", "-f"),
}

func init() {
	rootCmd.AddCommand(sourceCmd)
}
