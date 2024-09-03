package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	apiURL = "https://api.binance.com/api/v3/klines"
)

var (
	httpClient = &http.Client{Timeout: 10 * time.Second}
	marketDataCache struct {
		sync.RWMutex
		data map[string][]float64
		lastUpdate time.Time
	}
)

func init() {
	marketDataCache.data = make(map[string][]float64)
}

func fetchMarketData(symbol string, interval string, limit int) ([]float64, error) {
	fmt.Printf("Récupération des données du marché pour %s, intervalle %s, limite %d\n", symbol, interval, limit)

	cacheKey := fmt.Sprintf("%s-%s-%d", symbol, interval, limit)

	marketDataCache.RLock()
	if time.Since(marketDataCache.lastUpdate) < 10*time.Second {
		if cachedData, ok := marketDataCache.data[cacheKey]; ok {
			marketDataCache.RUnlock()
			return cachedData, nil
		}
	}
	marketDataCache.RUnlock()

	marketDataCache.Lock()
	defer marketDataCache.Unlock()

	if time.Since(marketDataCache.lastUpdate) < 10*time.Second {
		if cachedData, ok := marketDataCache.data[cacheKey]; ok {
			return cachedData, nil
		}
	}

	url := fmt.Sprintf("%s?symbol=%s&interval=%s&limit=%d", apiURL, symbol, interval, limit)
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var klines [][]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&klines); err != nil {
		return nil, err
	}

	closingPrices := make([]float64, 0, len(klines))
	for _, k := range klines {
		closePrice, _ := strconv.ParseFloat(k[4].(string), 64)
		closingPrices = append(closingPrices, closePrice)
	}

	return closingPrices, nil
}