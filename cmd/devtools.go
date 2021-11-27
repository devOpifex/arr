package cmd

import (
	"github.com/spf13/cobra"
)

var devtoolsCmd = &cobra.Command{
	Use:   "devtools",
	Short: "interact with the {devtools} package",
	Long:  `build, document, install, etc. packages`,
	Args:  cobra.MinimumNArgs(1),
	ValidArgs: []string{
		"document",
		"check",
		"check_rhub",
		"build",
		"install",
		"revdep",
		"test",
		"lint",
		"build_readme",
		"create",
		"submit_cran",
	},
	Run: makeRun("devtools", "-e"),
}

func init() {
	rootCmd.AddCommand(devtoolsCmd)
}
