package cmd

import (
	"github.com/spf13/cobra"
)

var devtoolsCmd = &cobra.Command{
	Use:   "devtools",
	Short: "interact with the {devtools} package",
	Long: `build, document, install, etc. packages.
Arguments are function names from the {devtools} package.

Examples:
	arr devtools document
	arr devtools check
	arr devtools document check install`,
	Args: cobra.MinimumNArgs(1),
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
	Run: makeRun("devtools", "-e -q"),
}

func init() {
	rootCmd.AddCommand(devtoolsCmd)
}
