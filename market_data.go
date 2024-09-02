package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const (
	apiURL = "https://api.binance.com/api/v3/klines"
)

func fetchMarketData(symbol string, interval string, limit int) ([]float64, error) {
	url := fmt.Sprintf("%s?symbol=%s&interval=%s&limit=%d", apiURL, symbol, interval, limit)
	fmt.Println("URL de la requête :", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var klines [][]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&klines); err != nil {
		return nil, err
	}

	var closingPrices []float64
	for _, k := range klines {
		closePrice, _ := strconv.ParseFloat(k[4].(string), 64)
		closingPrices = append(closingPrices, closePrice)
	}

	return closingPrices, nil
}