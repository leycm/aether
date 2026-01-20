package main

import (
	"fmt"
	"os"

	"aether/cli"
)

func main() {
	cfg, args := cli.ParseGlobal(os.Args[1:])

	if len(args) == 0 {
		cli.StartREPL(cfg)
		return
	}

	if err := cli.DispatchCommand(cfg, args); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}
