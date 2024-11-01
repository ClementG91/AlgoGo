package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func testBinanceConnection() error {
	resp, err := http.Get("https://api.binance.com/api/v3/ping")
	if err != nil {
		return fmt.Errorf("erreur de connexion à l'API Binance : %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("réponse inattendue de l'API Binance : %s", resp.Status)
	}
	return nil
}

func main() {
	if err := LoadConfig(); err != nil {
		handleError(err)
		return
	}
	fmt.Printf("Configuration chargée : %+v\n", AppConfig)
	fmt.Printf("Temps d'attente entre les cycles : %d secondes\n", AppConfig.SleepTime)

	fmt.Println("Création du ticker...")
	ticker := time.NewTicker(time.Duration(AppConfig.SleepTime) * time.Second)
	defer ticker.Stop()

	var wg sync.WaitGroup
	wg.Add(1)

	fmt.Println("Démarrage de la goroutine...")
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine démarrée, exécution du premier cycle de trading...")
		if err := runTradingCycle(); err != nil {
			fmt.Printf("Erreur dans le cycle de trading initial : %v\n", err)
			handleError(err)
			return
		}
		fmt.Println("En attente des prochains ticks...")
		for range ticker.C {
			fmt.Println("Tick reçu, démarrage d'un nouveau cycle de trading")
			if err := runTradingCycle(); err != nil {
				fmt.Printf("Erreur dans le cycle de trading : %v\n", err)
				handleError(err)
				return
			}
		}
	}()

	if err := testBinanceConnection(); err != nil {
		handleError(err)
		return
	}
	fmt.Println("Connexion à l'API Binance réussie")

	fmt.Println("En attente de la fin de la goroutine...")
	wg.Wait()
	fmt.Println("Programme terminé.")
}

// Structure pour gérer l'état du trading
type TradingState struct {
	mu sync.Mutex
}

var tradingState = TradingState{}

func runTradingCycle() error {
	fmt.Println("Début du cycle de trading")

	closingPrices, err := fetchMarketData(AppConfig.Symbol, AppConfig.Interval, 1000)
	if err != nil {
		return fmt.Errorf("erreur lors de la récupération des données du marché : %v", err)
	}
	if len(closingPrices) == 0 {
		return fmt.Errorf("aucun prix de clôture récupéré")
	}

	shortEMA := calculateEMA(closingPrices, AppConfig.ShortEMA)
	longEMA := calculateEMA(closingPrices, AppConfig.LongEMA)
	signal := generateSignal(shortEMA, longEMA)

	price := closingPrices[len(closingPrices)-1]
	
	// Calcul du PnL et enregistrement du trade si un signal est généré
	if signal != "HOLD" {
		tradingState.mu.Lock()
		
		if signal == "BUY" {
			GlobalPositionTracker.OpenPosition(
				AppConfig.Symbol,
				price,
				AppConfig.Quantity,
				shortEMA[len(shortEMA)-1],
				longEMA[len(longEMA)-1],
			)
		} else if signal == "SELL" {
			position := GlobalPositionTracker.ClosePosition(AppConfig.Symbol)
			if position != nil {
				trade := Trade{
					Timestamp:     time.Now(),
					Symbol:        AppConfig.Symbol,
					Type:         "SELL",
					Price:        price,
					Quantity:     position.Quantity,
					PnL:          (price - position.EntryPrice) * position.Quantity,
					PnLPercentage: ((price - position.EntryPrice) / position.EntryPrice) * 100,
				}
				
				if err := logTrade(trade); err != nil {
					tradingState.mu.Unlock()
					fmt.Printf("Erreur lors de l'enregistrement du trade : %v\n", err)
				}
			}
		}
		
		tradingState.mu.Unlock()
	}

	fmt.Printf("Signal généré : %s\n", signal)

	var orderErr error
	switch signal {
	case "BUY":
		fmt.Println("Tentative de placement d'un ordre d'achat")
		orderErr = placeOrder(AppConfig.Symbol, "BUY", "LIMIT", AppConfig.Quantity, price)
	case "SELL":
		fmt.Println("Tentative de placement d'un ordre de vente")
		orderErr = placeOrder(AppConfig.Symbol, "SELL", "LIMIT", AppConfig.Quantity, price)
	default:
		fmt.Println("Aucun ordre placé")
	}

	if orderErr != nil {
		return fmt.Errorf("erreur lors du placement de l'ordre : %v", orderErr)
	}

	if err := printAccountBalance(); err != nil {
		return fmt.Errorf("erreur lors de l'affichage du solde du compte : %v", err)
	}

	fmt.Println("Fin du cycle de trading")
	return nil
}