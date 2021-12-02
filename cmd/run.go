package cmd

import (
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run code",
	Long: `run arbitrary functions or expressions
Example: arr run 'packer::bundle()'`,
	Args: cobra.MinimumNArgs(1),
	Run:  makeRun("", "-e -q"),
}

func init() {
	rootCmd.AddCommand(runCmd)
}
