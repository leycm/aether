package commands

import (
	"aether/config"
	"aether/constants"
	"flag"
	"fmt"
	"strings"
)

// ModelSize represents the size of the model to train
type ModelSize string

// Supported model sizes
const (
	ModelSizeSmall  ModelSize = "s"
	ModelSizeMedium ModelSize = "m"
	ModelSizeLarge  ModelSize = "l"
)

// Train handles the training of prediction models
func Train(cfg *config.Config, args []string) error {
	fs := flag.NewFlagSet("train", flag.ContinueOnError)
	fs.Usage = func() {
		fmt.Fprintf(
			flag.CommandLine.Output(),
			"Usage: %s [options]\n\nOptions:\n",
			fs.Name(),
		)
		fs.PrintDefaults()
	}

	// Define flags
	size := fs.String(
		"size",
		string(ModelSizeMedium),
		fmt.Sprintf(
			"model size (%s, %s, %s)",
			ModelSizeSmall,
			ModelSizeMedium,
			ModelSizeLarge,
		),
	)

	crypto := fs.Bool(
		"crypto",
		false,
		"train model on cryptocurrency data",
	)
	stock := fs.Bool(
		"stock",
		true,
		"train model on stock market data",
	)
	etf := fs.Bool(
		"etf",
		true,
		"train model on ETF data",
	)

	// Parse flags
	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("%w: %v", constants.ErrInvalidArgs, err)
	}

	// Validate model size
	modelSize := ModelSize(strings.ToLower(*size))
	switch modelSize {
	case ModelSizeSmall, ModelSizeMedium, ModelSizeLarge:
		// Valid size
	default:
		return fmt.Errorf(
			"%w: invalid model size '%s', must be one of: %s, %s, %s",
			constants.ErrInvalidArgs,
			*size,
			ModelSizeSmall,
			ModelSizeMedium,
			ModelSizeLarge,
		)
	}

	// Validate at least one data type is selected
	if !*crypto && !*stock && !*etf {
		return fmt.Errorf(
			"%w: at least one data type must be selected (use --crypto, --stock, or --etf)",
			constants.ErrInvalidArgs,
		)
	}

	// Log training parameters
	logTrainingParams(cfg.Workspace, modelSize, *crypto, *stock, *etf)

	// TODO: Implement model training
	// 1. Load and preprocess training data
	// 2. Initialize model with specified size
	// 3. Train model on selected data types
	// 4. Save trained model to workspace

	return nil
}

// logTrainingParams logs the training parameters in a consistent format
func logTrainingParams(workspace string, size ModelSize, crypto, stock, etf bool) {
	const (
		header = "TRAINING PARAMETERS"
		format = " %-12s: %v\n"
	)

	fmt.Println("\n" + header)
	fmt.Printf(format, "Workspace", workspace)
	fmt.Printf(format, "Model Size", strings.ToUpper(string(size)))
	fmt.Printf(format, "Crypto", crypto)
	fmt.Printf(format, "Stocks", stock)
	fmt.Printf(format, "ETFs", etf)
	fmt.Println()
}
