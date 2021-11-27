package cmd

import (
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run code",
	Long:  `Run arbitrary functions or expressions`,
	Args:  cobra.MinimumNArgs(1),
	Run:   makeRun("", "-e"),
}

func init() {
	rootCmd.AddCommand(runCmd)
}
