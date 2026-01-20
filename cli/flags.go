package cli

import (
	"aether/config"
	"flag"
)

func ParseGlobal(args []string) (*config.Config, []string) {
	fs := flag.NewFlagSet("aether", flag.ContinueOnError)

	workspace := fs.String("w", "./", "workspace directory")
	fs.StringVar(workspace, "workspace", "./", "workspace directory")

	_ = fs.Parse(args)

	return &config.Config{
		Workspace: *workspace,
	}, fs.Args()
}
