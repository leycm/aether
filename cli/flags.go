package cli

import (
	"aether/config"
	"aether/constants"
	"flag"
	"fmt"
)

// ParseGlobal parses the global command-line flags and returns the configuration and remaining arguments
func ParseGlobal(args []string) (*config.Config, []string, error) {
	fs := flag.NewFlagSet("aether", flag.ContinueOnError)

	// Define flags
	workspace := fs.String(
		"w",
		"./",
		"workspace directory (default: current directory)",
	)
	fs.StringVar(
		workspace,
		"workspace",
		"./",
		"workspace directory (default: current directory)",
	)

	// Parse flags
	if err := fs.Parse(args); err != nil {
		return nil, nil, fmt.Errorf("%w: %v", constants.ErrInvalidArgs, err)
	}

	// Validate workspace path
	if *workspace == "" {
		return nil, nil, fmt.Errorf("%w: workspace path cannot be empty", constants.ErrInvalidArgs)
	}

	return &config.Config{
		Workspace: *workspace,
	}, fs.Args(), nil
}
