package cli

import (
	"aether/constants"
	"fmt"
	"strings"

	"aether/config"

	"github.com/chzyer/readline"
)

// REPL commands
const (
	CmdExit = "exit"
	CmdQuit = "quit"
)

// StartREPL starts the interactive Read-Eval-Print Loop
func StartREPL(cfg *config.Config) error {
	// Initialize readline with tab completion
	completer := readline.NewPrefixCompleter(
		readline.PcItem(constants.CmdTrain),
		readline.PcItem(constants.CmdDownload),
		readline.PcItem(constants.CmdPredict),
		readline.PcItem(CmdExit),
		readline.PcItem(CmdQuit),
	)

	// Configure readline
	rl, err := readline.NewEx(&readline.Config{
		Prompt:            "> ",
		HistoryFile:       cfg.Workspace + "/.aether_history",
		InterruptPrompt:   "^C",
		EOFPrompt:         CmdExit,
		HistorySearchFold: true,
		AutoComplete:      completer,
	})
	if err != nil {
		return fmt.Errorf("%w: failed to initialize readline: %v", constants.ErrConfigLoad, err)
	}
	defer rl.Close()

	fmt.Println("Aether interactive console (type 'exit' or 'quit' to exit)")

	// Main REPL loop
	for {
		line, err := rl.Readline()
		if err != nil {
			if err == readline.ErrInterrupt {
				fmt.Println("\nUse 'exit' or 'quit' to exit")
				continue
			}
			if err.Error() == "EOF" {
				break
			}
			return fmt.Errorf("%w: readline error: %v", constants.ErrConfigLoad, err)
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		switch line {
		case CmdExit, CmdQuit:
			return nil
		}

		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}

		if err := DispatchCommand(cfg, args); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}

	return nil
}
