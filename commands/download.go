package commands

import (
	"aether/config"
	load "aether/download"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func Download(cfg *config.Config, args []string) error {
	fs := flag.NewFlagSet("download", flag.ContinueOnError)

	crypto := fs.Bool("C", false, "download crypto data")
	stock := fs.Bool("S", false, "download stock data")
	etf := fs.Bool("E", false, "download etf data")

	_ = fs.Parse(args)

	fmt.Println("DOWNLOAD")
	fmt.Println(" workspace:", cfg.Workspace)
	fmt.Println(" crypto:", *crypto)
	fmt.Println(" stock :", *stock)
	fmt.Println(" etf   :", *etf)

	resources := filepath.Join(cfg.Workspace, ".resources")
	if err := os.MkdirAll(resources, os.ModePerm); err != nil {
		return fmt.Errorf("fail to create resources folder: %w", err)
	}

	if *crypto || *stock || *etf {

		if *stock {
			load.FetchAndSave("most_actives", filepath.Join(resources, "stock.symbols"))
		}
		if *etf {
			load.FetchAndSave("most_actives_etfs", filepath.Join(resources, "etf.symbols"))
		}
		if *crypto {
			load.FetchAndSave("all_cryptocurrencies_us", filepath.Join(resources, "crypto.symbols"))
		}
	}

	// TODO: Rate-Limit / Retry-Handling
	// TODO: Daten lokal speichern
	// TODO: Fortschritt + Logs ausgeben

	return nil
}
