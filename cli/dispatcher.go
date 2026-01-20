package cli

import (
	"aether/config"
	"errors"

	"aether/commands"
)

func DispatchCommand(cfg *config.Config, args []string) error {
	if len(args) == 0 {
		return errors.New("no commands")
	}

	cmd := args[0]
	cmdArgs := args[1:]

	switch cmd {
	case "train":
		return commands.Train(cfg, cmdArgs)
	case "download":
		return commands.Download(cfg, cmdArgs)
	case "predict":
		return commands.Predict(cfg, cmdArgs)
	default:
		return errors.New("unknown commands: " + cmd)
	}
}
