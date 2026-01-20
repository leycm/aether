package cli

import (
	"errors"
	"fmt"
	"strings"

	"aether/config"

	"github.com/chzyer/readline"
)

func StartREPL(cfg *config.Config) {
	completer := readline.NewPrefixCompleter(
		readline.PcItem("train"),
		readline.PcItem("download"),
		readline.PcItem("predict"),
		readline.PcItem("exit"),
		readline.PcItem("quit"),
	)

	rl, err := readline.NewEx(&readline.Config{
		Prompt:            "> ",
		HistoryFile:       cfg.Workspace + "/.history",
		InterruptPrompt:   "^C",
		EOFPrompt:         "exit",
		HistorySearchFold: true,
		AutoComplete:      completer,
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	fmt.Println("Aether interactive console (type 'exit' to quit)")

	for {
		line, err := rl.Readline()
		if errors.Is(err, readline.ErrInterrupt) {
			break
		}
		if err != nil {
			break
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if line == "exit" || line == "quit" {
			break
		}

		args := strings.Fields(line)
		if err := DispatchCommand(cfg, args); err != nil {
			fmt.Println("Error:", err)
		}
	}
}
