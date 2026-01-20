package commands

import (
	"aether/config"
	"flag"
	"fmt"
)

func Train(cfg *config.Config, args []string) error {
	fs := flag.NewFlagSet("train", flag.ContinueOnError)

	size := fs.String("s", "m", "model size (s,m,l)")
	fs.StringVar(size, "size", "m", "model size")

	crypto := fs.Bool("C", false, "crypto")
	stock := fs.Bool("S", false, "stock")
	etf := fs.Bool("E", false, "etf")

	_ = fs.Parse(args)

	fmt.Println("TRAIN")
	fmt.Println(" workspace:", cfg.Workspace)
	fmt.Println(" size:", *size)
	fmt.Println(" crypto:", *crypto, "stock:", *stock, "etf:", *etf)

	// TODO: Training starten
	return nil
}
