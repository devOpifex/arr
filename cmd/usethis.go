package cmd

import (
	"github.com/spf13/cobra"
)

var usethisCmd = &cobra.Command{
	Use:   "usethis",
	Short: "interact with the {usethis} package",
	Long: `create packages, website, etc.
Examples:
	arr usethis use_gpl_license -1=3
	arr usethis create_package --arg1s=myPkg
`,
	Args: cobra.MinimumNArgs(1),
	ValidArgs: []string{
		"use_build_ignore",
		"use_git_ignore",
		"use_make",
		"use_mit_license",
		"use_gpl_license",
		"use_gpl3_license",
		"use_agpl_license",
		"use_apache_license",
		"use_data",
		"use_rcpp",
		"use_cpp11",
		"use_pipe",
		"use_test",
		"use_pkgdown",
		"use_appveyor",
		"use_jenkins",
		"use_travis",
		"use_rstudio",
		"use_git",
		"use_github",
		"create_package",
	},
	Run: makeRun("usethis", "-e"),
}

func init() {
	rootCmd.AddCommand(usethisCmd)
}
