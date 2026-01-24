package cli

import (
	"aether/commands"
	"aether/config"
	"aether/constants"
	"fmt"
)

// DispatchCommand routes the command to the appropriate handler
func DispatchCommand(cfg *config.Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf(constants.ErrNoCommands)
	}

	cmd := args[0]
	cmdArgs := args[1:]

	switch cmd {
	case constants.CmdTrain:
		if err := commands.Train(cfg, cmdArgs); err != nil {
			return fmt.Errorf("train command failed: %w", err)
		}
	case constants.CmdDownload:
		if err := commands.Download(cfg, cmdArgs); err != nil {
			return fmt.Errorf("download command failed: %w", err)
		}
	case constants.CmdPredict:
		if err := commands.Predict(cfg, cmdArgs); err != nil {
			return fmt.Errorf("predict command failed: %w", err)
		}
	default:
		return fmt.Errorf(constants.ErrUnknownCommand, cmd)
	}

	return nil
}
