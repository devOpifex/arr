package cmd

import (
	"github.com/spf13/cobra"
)

var sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Interact with the {devtools} package",
	Long:  `Build, document, install, etc. packages`,
	Args:  cobra.MinimumNArgs(1),
	Run:   makeRun("", true),
}

func init() {
	rootCmd.AddCommand(sourceCmd)
}
