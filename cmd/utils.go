package cmd

import (
	"bufio"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

func getPath() string {
	var p string

	if path != "" {
		return path
	}

	p, err := exec.LookPath("R")

	if err != nil {
		fmt.Println("Could not locate R installation")
	}

	return p
}

func makeArg(pkg, arg, first string) string {
	if pkg != "" {
		return pkg + "::" + arg + "();"
	}

	if first == "-f" {
		return arg
	}

	return arg + ";"
}

func makeRun(pkg, first string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		rArgs := first + " "

		for _, arg := range args {
			rArgs += makeArg(pkg, arg, first)
		}

		path := getPath()

		rCommand := exec.Command(path, rArgs)
		stdout, _ := rCommand.StdoutPipe()

		rCommand.Start()
		if verbose {
			scanner := bufio.NewScanner(stdout)
			scanner.Split(bufio.ScanLines)
			for scanner.Scan() {
				line := scanner.Text()
				fmt.Println(line)
			}
		}
		rCommand.Wait()
	}
}
