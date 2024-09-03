package main

import "fmt"

func generateSignal(shortEMA, longEMA []float64) string {
	fmt.Println("Génération du signal de trading")
	if shortEMA[len(shortEMA)-1] > longEMA[len(longEMA)-1] && shortEMA[len(shortEMA)-2] <= longEMA[len(longEMA)-2] {
		return "BUY"
	} else if shortEMA[len(shortEMA)-1] < longEMA[len(longEMA)-1] && shortEMA[len(shortEMA)-2] >= longEMA[len(longEMA)-2] {
		return "SELL"
	}
	return "HOLD"
}