package cmd

import (
	"bufio"
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
		path, err := exec.LookPath("R")

		if err != nil {
			fmt.Println("Could not locate R installation")
		}

		rCommand := exec.Command(path, rArgs)
		stdout, _ := rCommand.StdoutPipe()

		rCommand.Start()
		scanner := bufio.NewScanner(stdout)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}
		rCommand.Wait()
	}
}
