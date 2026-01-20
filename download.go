package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"time"

	yfa "github.com/oscarli916/yahoo-finance-api"
)

func downloadTop500() {
	fmt.Println("Downloading top500.csv...")
	file, err := os.Open("top500.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = ';'
	tickers, _ := r.ReadAll()

	os.MkdirAll(".data", os.ModePerm)

	for _, line := range tickers {
		ticker := line[1]
		filename := ".data/" + ticker + ".json"

		if _, err := os.Stat(filename); err == nil {
			fmt.Println(ticker, "bereits gespeichert, überspringe...")
			continue
		}

		fmt.Println("Lade [", line[0], "]:", ticker, " with value of ", line[2])
		t := yfa.NewTicker(ticker)
		history, err := t.History(yfa.HistoryQuery{
			Range:    "2y",
			Interval: "1h",
		})
		if err != nil {
			fmt.Println("Fehler bei", ticker, ":", err)
			continue
		}

		f, err := os.Create(filename)
		if err != nil {
			fmt.Println("Fehler beim Erstellen der Datei:", err)
			continue
		}
		enc := json.NewEncoder(f)
		enc.SetIndent("", "  ")
		enc.Encode(history)
		f.Close()

		time.Sleep(time.Second)
	}
}
