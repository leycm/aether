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
			file := filepath.Join(resources, "stock.symbols")
			destination := filepath.Join(resources, "stock")
			load.FetchAndSave("most_actives", file)
			load.DownloadHistory(file, destination)
		}
		if *etf {
			file := filepath.Join(resources, "etf.symbols")
			destination := filepath.Join(resources, "etf")
			load.FetchAndSave("most_actives_etfs", file)
			load.DownloadHistory(file, destination)
		}
		if *crypto {
			file := filepath.Join(resources, "crypto.symbols")
			destination := filepath.Join(resources, "crypto")
			load.FetchAndSave("all_cryptocurrencies_us", file)
			load.DownloadHistory(file, destination)
		}
	}

	return nil
}
