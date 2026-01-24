package main

import (
	"fmt"
	"os"

	"aether/cli"
	"aether/constants"
)

func main() {
	cfg, args, err := cli.ParseGlobal(os.Args[1:])
	if err != nil {
		exitWithError(err, constants.ExitConfigError)
	}

	if len(args) == 0 {
		if err := cli.StartREPL(cfg); err != nil {
			exitWithError(err, constants.ExitError)
		}
		return
	}

	if err := cli.DispatchCommand(cfg, args); err != nil {
		exitWithError(err, constants.ExitError)
	}
}

func exitWithError(err error, code int) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	os.Exit(code)
}
