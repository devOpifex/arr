package cmd

import (
	"github.com/spf13/cobra"
)

var usethisCmd = &cobra.Command{
	Use:   "usethis",
	Short: "Interact with the {usethis} package",
	Long:  `Call use_* functions.`,
	Args:  cobra.MinimumNArgs(1),
	ValidArgs: []string{
		"build_ignore",
		"git_ignore",
		"make",
		"mit_license",
		"gpl_license",
		"gpl3_license",
		"agpl_license",
		"apache_license",
		"data",
		"rcpp",
		"cpp11",
		"pipe",
		"test",
		"pkgdown",
		"appveyor",
		"jenkins",
		"travis",
		"rstudio",
		"git",
		"github",
	},
	Run: makeRun("usethis"),
}

func init() {
	rootCmd.AddCommand(usethisCmd)
}
