package commands

import (
	"aether/config"
	"aether/constants"
	load "aether/download"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// Download handles the download command for fetching market data
func Download(cfg *config.Config, args []string) error {
	fs := flag.NewFlagSet("download", flag.ContinueOnError)
	fs.Usage = func() {
		fmt.Fprintf(
			flag.CommandLine.Output(),
			"Usage: %s [options]\nOptions:\n",
			fs.Name(),
		)
		fs.PrintDefaults()
	}

	// Define flags
	crypto := fs.Bool(
		"crypto",
		false,
		"download cryptocurrency market data",
	)
	stock := fs.Bool(
		"stock",
		false,
		"download stock market data",
	)
	etf := fs.Bool(
		"etf",
		false,
		"download ETF market data",
	)

	// Parse flags
	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("%w: %v", constants.ErrInvalidArgs, err)
	}

	// Validate at least one data type is selected
	if !*crypto && !*stock && !*etf {
		fs.Usage()
		return fmt.Errorf("%w: no data type specified (use -crypto, -stock, or -etf)", constants.ErrInvalidArgs)
	}

	// Create resources directory
	resources := filepath.Join(cfg.Workspace, ".resources")
	if err := os.MkdirAll(resources, 0755); err != nil {
		return fmt.Errorf("%w: failed to create resources directory: %v", constants.ErrFileOperation, err)
	}

	// Process downloads
	downloads := []struct {
		enabled bool
		name    string
		source  string
	}{
		{*stock, "stock", "most_actives"},
		{*etf, "etf", "most_actives_etfs"},
		{*crypto, "crypto", "all_cryptocurrencies_us"},
	}

	for _, d := range downloads {
		if !d.enabled {
			continue
		}

		symbolsFile := filepath.Join(resources, d.name+".symbols")
		destination := filepath.Join(resources, d.name)

		fmt.Printf("Downloading %s symbols...\n", d.name)
		if err := load.FetchAndSave(d.source, symbolsFile); err != nil {
			return fmt.Errorf("failed to fetch %s symbols: %w", d.name, err)
		}

		fmt.Printf("Downloading %s history...\n", d.name)
		if err := load.DownloadHistory(symbolsFile, destination); err != nil {
			return fmt.Errorf("failed to download %s history: %w", d.name, err)
		}

		fmt.Printf("Successfully downloaded %s data\n", d.name)
	}

	return nil
}
