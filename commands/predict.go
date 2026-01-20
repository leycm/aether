package commands

import (
	"aether/config"
	"aether/util"
	"errors"
	"flag"
	"fmt"
	"time"
)

func Predict(cfg *config.Config, args []string) error {
	if len(args) == 0 {
		return errors.New("predict requires <time>")
	}

	targetTimeRaw := args[0]
	cmdArgs := args[1:]

	fs := flag.NewFlagSet("predict", flag.ContinueOnError)

	step := fs.Duration("s", time.Hour, "step size (e.g. 1h)")
	fs.DurationVar(step, "step", time.Hour, "step size")

	crypto := fs.Bool("C", false, "predict crypto")
	stock := fs.Bool("S", false, "predict stock")
	etf := fs.Bool("E", false, "predict etf")

	_ = fs.Parse(cmdArgs)

	// Default: alles
	if !*crypto && !*stock && !*etf {
		*crypto, *stock, *etf = true, true, true
	}

	targetTime, err := util.ParseFutureTime(targetTimeRaw)
	if err != nil {
		return err
	}

	fmt.Println("PREDICT")
	fmt.Println(" workspace:", cfg.Workspace)
	fmt.Println(" target   :", targetTime)
	fmt.Println(" step     :", *step)
	fmt.Println(" crypto   :", *crypto)
	fmt.Println(" stock    :", *stock)
	fmt.Println(" etf      :", *etf)

	// TODO: Lade trainierte Modelle aus Workspace
	// TODO: Prüfe ob TargetTime > now
	// TODO: Erzeuge Zeitachse (now → target, step)
	// TODO: Führe Inference pro Step aus
	// TODO: Aggregiere + speichere Predictions
	// TODO: Ausgabe (Console / JSON / CSV)

	return nil
}
