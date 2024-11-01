package main

import "fmt"

func calculateEMA(prices []float64, period int) []float64 {
	fmt.Printf("Calcul de l'EMA pour une période de %d\n", period)
	ema := make([]float64, len(prices))
	multiplier := 2.0 / float64(period+1)
	ema[0] = prices[0]

	for i := 1; i < len(prices); i++ {
		ema[i] = ((prices[i] - ema[i-1]) * multiplier) + ema[i-1]
	}

	fmt.Printf("Dernière valeur EMA-%d : %f\n", period, ema[len(ema)-1])
	fmt.Printf("Avant-dernière valeur EMA-%d : %f\n", period, ema[len(ema)-2])

	return ema
}