package cmd

import (
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Interact with the {devtools} package",
	Long:  `Build, document, install, etc. packages`,
	Args:  cobra.MinimumNArgs(1),
	Run:   makeRun("", false),
}

func init() {
	rootCmd.AddCommand(runCmd)
}
