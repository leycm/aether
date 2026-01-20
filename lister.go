package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	baseURL = "https://query1.finance.yahoo.com/v1/finance/screener/predefined/saved"
	step    = 250
	delay   = 1 * time.Second
)

type Response struct {
	Finance struct {
		Result []struct {
			Quotes []struct {
				Symbol string `json:"symbol"`
			} `json:"quotes"`
		} `json:"result"`
	} `json:"finance"`
}

func list() {
	fetchAndSave("most_actives", "lists/stock_symbols.list")
	fetchAndSave("most_actives_etfs", "lists/etf_symbols.list")
	fetchAndSave("all_cryptocurrencies_us", "lists/crypto_symbols.list")
}

func fetchAndSave(scrId, filename string) {
	fmt.Println("Starte:", scrId)

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	client := &http.Client{}
	total := 0

	for start := 0; ; start += step {
		url := fmt.Sprintf(
			"%s?scrIds=%s&count=%d&start=%d",
			baseURL, scrId, step, start,
		)

		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("User-Agent", "Mozilla/5.0")
		req.Header.Set("Accept", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode != 200 {
			resp.Body.Close()
			log.Fatalf("[%s] HTTP %d", scrId, resp.StatusCode)
		}

		var data Response
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			resp.Body.Close()
			log.Fatal(err)
		}
		resp.Body.Close()

		if len(data.Finance.Result) == 0 ||
			len(data.Finance.Result[0].Quotes) == 0 {
			fmt.Println("Fertig:", scrId)
			break
		}

		for _, q := range data.Finance.Result[0].Quotes {
			file.WriteString(q.Symbol + "\n")
			total++
		}

		fmt.Printf("[%s] start=%d (+%d)\n", scrId, start, len(data.Finance.Result[0].Quotes))
		time.Sleep(delay)
	}

	fmt.Printf("[%s] Gespeichert: %d Symbole → %s\n\n", scrId, total, filename)
}
