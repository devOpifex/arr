package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var devtoolsCmd = &cobra.Command{
	Use:   "devtools",
	Short: "Interact with the {devtools} package",
	Long:  `Build, document, install, etc. packages`,
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
	Run: func(cmd *cobra.Command, args []string) {
		rArgs := "-e "
		for _, arg := range args {
			rArgs += "devtools::" + arg + "();"
		}
		fmt.Println(rArgs)
		path, _ := exec.LookPath("R")
		rCommand := exec.Command(path, rArgs)

		if err := rCommand.Run(); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(devtoolsCmd)
}
