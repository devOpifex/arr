package cmd

import "github.com/spf13/cobra"

var rsconnectCmd = &cobra.Command{
	Use:   "rsconnect",
	Short: "interact with the {rsconnect} package",
	Long:  `call use_* functions.`,
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
