package main

func generateSignal(shortEMA, longEMA []float64) string {
	if shortEMA[len(shortEMA)-1] > longEMA[len(longEMA)-1] && shortEMA[len(shortEMA)-2] <= longEMA[len(longEMA)-2] {
		return "BUY"
	} else if shortEMA[len(shortEMA)-1] < longEMA[len(longEMA)-1] && shortEMA[len(shortEMA)-2] >= longEMA[len(longEMA)-2] {
		return "SELL"
	}
	return "HOLD"
}