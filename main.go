package main

import (
	"fmt"
	"time"
)

func main() {
	err := LoadConfig()
	if err != nil {
		handleError(err)
		return
	}

	for {
		closingPrices, err := fetchMarketData(AppConfig.Symbol, AppConfig.Interval, 100)
		if err != nil {
			handleError(err)
			continue
		}

		shortEMA := calculateEMA(closingPrices, AppConfig.ShortEMA)
		longEMA := calculateEMA(closingPrices, AppConfig.LongEMA)
		signal := generateSignal(shortEMA, longEMA)

		fmt.Println("Signal :", signal)

		price := closingPrices[len(closingPrices)-1]

		switch signal {
		case "BUY":
			err = placeOrder(AppConfig.Symbol, "BUY", "LIMIT", AppConfig.Quantity, price)
		case "SELL":
			err = placeOrder(AppConfig.Symbol, "SELL", "LIMIT", AppConfig.Quantity, price)
		}

		if err != nil {
			handleError(err)
		}

		printAccountBalance()

		time.Sleep(time.Duration(AppConfig.SleepTime) * time.Second)
	}
}