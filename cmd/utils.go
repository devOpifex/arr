package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

func makeRun(pkg string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		rArgs := "-e "
		for _, arg := range args {
			rArgs += pkg + "::" + arg + "();"
		}
		path, _ := exec.LookPath("R")
		rCommand := exec.Command(path, rArgs)

		if err := rCommand.Run(); err != nil {
			fmt.Println(err)
		}
	}
}
