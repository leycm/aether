package dev

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"time"

	yfa "github.com/oscarli916/yahoo-finance-api"
)

func DownloadHistory(filename string, destination string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = ';'
	tickers, _ := r.ReadAll()

	err = os.MkdirAll(destination, os.ModePerm)
	if err != nil {
		return
	}

	for _, line := range tickers {
		ticker := line[0]
		filename := destination + "/" + ticker

		if _, err := os.Stat(filename); err == nil {
			fmt.Println(ticker, "downloaded, skip...")
			continue
		}

		fmt.Println("Load:", ticker)
		t := yfa.NewTicker(ticker)
		history, err := t.History(yfa.HistoryQuery{
			Range:    "2y",
			Interval: "1h",
		})
		if err != nil {
			fmt.Println("Error on", ticker, ":", err)
			continue
		}

		f, err := os.Create(filename)
		if err != nil {
			fmt.Println("Error while creating the File:", err)
			continue
		}
		enc := json.NewEncoder(f)
		enc.SetIndent("", "  ")
		enc.Encode(history)
		f.Close()

		time.Sleep(time.Second)
	}
}
