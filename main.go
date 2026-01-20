package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	list()
}

func oldmain() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "download":
		handleDownload()
	case "train":
		handleTrain()
	case "predict":
		handlePredict()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  download [uritoreadfrom|.data/top500.csv] [pathtodest|.data]")
	fmt.Println("  train [datapath|.data] [name|default.model]")
	fmt.Println("  predict <ticker> <time> [time|1h] [model|default.model] [datapath|.data]")
	fmt.Println("\nExample:")
	fmt.Println("  predict GOOGL 24h -t 1h -d .data2")
}

func handleDownload() {
	cmd := flag.NewFlagSet("download", flag.ExitOnError)

	uriPtr := cmd.String("list", ".data/top500.csv", "Source URI to download from")
	destPtr := cmd.String("dest", ".data", "Destination path")

	cmd.Parse(os.Args[2:])

	fmt.Printf("Download command called\n")
	fmt.Printf("  URI: %s\n", *uriPtr)
	fmt.Printf("  Destination: %s\n", *destPtr)

	// downloadData(*uriPtr, *destPtr)
}

func handleTrain() {
	cmd := flag.NewFlagSet("train", flag.ExitOnError)

	dataPathPtr := cmd.String("data", ".data", "Path to training data")
	namePtr := cmd.String("name", "default.model", "Model name")

	cmd.Parse(os.Args[2:])

	fmt.Printf("Train command called\n")
	fmt.Printf("  Data path: %s\n", *dataPathPtr)
	fmt.Printf("  Model name: %s\n", *namePtr)

	// trainModel(*dataPathPtr, *namePtr)
}

func handlePredict() {
	if len(os.Args) < 4 {
		fmt.Println("Error: ticker and time are required for predict command")
		fmt.Println("Usage: predict <ticker> <time> [time|1h] [model|default.model] [datapath|.data]")
		os.Exit(1)
	}

	cmd := flag.NewFlagSet("predict", flag.ExitOnError)

	ticker := os.Args[2]
	timeArg := os.Args[3]

	timeResPtr := cmd.String("t", "1h", "Time resolution (cannot be lower than 1h)")
	modelPtr := cmd.String("m", "default.model", "Model to use for prediction")
	dataPathPtr := cmd.String("d", ".data", "Path to data directory")

	cmd.StringVar(timeResPtr, "time", "1h", "Time resolution (cannot be lower than 1h)")
	cmd.StringVar(modelPtr, "model", "default.model", "Model to use for prediction")
	cmd.StringVar(dataPathPtr, "datapath", ".data", "Path to data directory")

	cmd.Parse(os.Args[4:])

	if *timeResPtr < "1h" {
		fmt.Printf("Error: Time resolution cannot be lower than 1h, got %s\n", *timeResPtr)
		os.Exit(1)
	}

	fmt.Printf("Predict command called\n")
	fmt.Printf("  Ticker: %s\n", ticker)
	fmt.Printf("  Time: %s\n", timeArg)
	fmt.Printf("  Time resolution: %s\n", *timeResPtr)
	fmt.Printf("  Model: %s\n", *modelPtr)
	fmt.Printf("  Data path: %s\n", *dataPathPtr)

	// predict(ticker, timeArg, *timeResPtr, *modelPtr, *dataPathPtr)
}
