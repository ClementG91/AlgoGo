package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type Trade struct {
	Timestamp     time.Time
	Symbol        string
	Type          string // "BUY" ou "SELL"
	Price         float64
	Quantity      float64
	PnL           float64  // Profit and Loss
	PnLPercentage float64
}

func logTrade(trade Trade) error {
	filename := "trades.csv"
	var file *os.File

	// Vérifie si le fichier existe
	if _, statErr := os.Stat(filename); os.IsNotExist(statErr) {
		var createErr error
		file, createErr = os.Create(filename)
		if createErr != nil {
			return fmt.Errorf("erreur lors de la création du fichier : %v", createErr)
		}
		writer := csv.NewWriter(file)
		// Écriture des en-têtes
		headers := []string{
			"Timestamp", "Symbol", "Type", "Price", "Quantity",
			"PnL", "PnLPercentage",
		}
		if err := writer.Write(headers); err != nil {
			return fmt.Errorf("erreur lors de l'écriture des en-têtes : %v", err)
		}
		writer.Flush()
	} else {
		var openErr error
		file, openErr = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
		if openErr != nil {
			return fmt.Errorf("erreur lors de l'ouverture du fichier : %v", openErr)
		}
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Conversion des données en strings
	record := []string{
		trade.Timestamp.Format(time.RFC3339),
		trade.Symbol,
		trade.Type,
		fmt.Sprintf("%.2f", trade.Price),
		fmt.Sprintf("%.8f", trade.Quantity),
		fmt.Sprintf("%.2f", trade.PnL),
		fmt.Sprintf("%.2f", trade.PnLPercentage),
	}

	if err := writer.Write(record); err != nil {
		return fmt.Errorf("erreur lors de l'écriture du trade : %v", err)
	}

	fmt.Printf("Trade enregistré : %+v\n", trade)
	return nil
} 