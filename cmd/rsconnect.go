package cmd

import "github.com/spf13/cobra"

var rsconnectCmd = &cobra.Command{
	Use:   "usethis",
	Short: "Interact with the {usethis} package",
	Long:  `Call use_* functions.`,
	Args:  cobra.MinimumNArgs(1),
	ValidArgs: []string{
		"writeManifest",
		"terminateApp",
		"lint",
		"accountUsage",
		"deployApp",
	},
	Run: makeRun("rsconnect", "-e"),
}

func init() {
	rootCmd.AddCommand(rsconnectCmd)
}
