package main

import "fmt"

func generateSignal(shortEMA, longEMA []float64) string {
	fmt.Println("Génération du signal de trading")
	fmt.Printf("Croisement EMA - Valeurs actuelles : Court %.2f / Long %.2f\n", 
		shortEMA[len(shortEMA)-1], longEMA[len(longEMA)-1])
	fmt.Printf("Croisement EMA - Valeurs précédentes : Court %.2f / Long %.2f\n", 
		shortEMA[len(shortEMA)-2], longEMA[len(longEMA)-2])

	if shortEMA[len(shortEMA)-1] > longEMA[len(longEMA)-1] && shortEMA[len(shortEMA)-2] <= longEMA[len(longEMA)-2] {
		return "BUY"
	} else if shortEMA[len(shortEMA)-1] < longEMA[len(longEMA)-1] && shortEMA[len(shortEMA)-2] >= longEMA[len(longEMA)-2] {
		return "SELL"
	}
	return "HOLD"
}