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

func makeArgs() string {
	var args string

	if arg1 == "" && arg1s == "" {
		return args
	}

	if arg1 != "" {
		args += arg1
	}

	if arg1s != "" {
		args += "'" + arg1s + "'"
	}

	if arg2 == "" && arg2s == "" {
		return args
	}

	args += ","
	if arg2 != "" {
		args += arg2
	}

	if arg2s != "" {
		args += "'" + arg2s + "'"
	}

	if arg3 == "" && arg3s == "" {
		return args
	}

	args += ","
	if arg3 != "" {
		args += arg3
	}

	if arg3s != "" {
		args += "'" + arg3s + "'"
	}

	return args
}

func makeCall(pkg, expr, first string) string {
	if pkg != "" {
		return pkg + "::" + expr + "(" + makeArgs() + ");"
	}

	if first == "-f" {
		return expr
	}

	return expr + ";"
}

func makeRun(pkg, first string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		var rArgs string

		for _, arg := range args {
			rArgs += makeCall(pkg, arg, first)
		}

		path := getPath()

		rCommand := exec.Command(path, first, rArgs)
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
