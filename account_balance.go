package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

const (
	accountInfoURL = "https://testnet.binance.vision/api/v3/account"
)

type AccountInfo struct {
	Balances []struct {
		Asset  string `json:"asset"`
		Free   string `json:"free"`
		Locked string `json:"locked"`
	} `json:"balances"`
}

var (
	accountInfoCache struct {
		sync.RWMutex
		info *AccountInfo
		lastUpdate time.Time
	}
)

func getAccountInfo() (*AccountInfo, error) {
	accountInfoCache.RLock()
	if time.Since(accountInfoCache.lastUpdate) < 5*time.Second {
		defer accountInfoCache.RUnlock()
		return accountInfoCache.info, nil
	}
	accountInfoCache.RUnlock()

	accountInfoCache.Lock()
	defer accountInfoCache.Unlock()

	if time.Since(accountInfoCache.lastUpdate) < 5*time.Second {
		return accountInfoCache.info, nil
	}

	timestamp := time.Now().Unix() * 1000
	params := fmt.Sprintf("timestamp=%d", timestamp)
	signature := signRequest(params)
	params = fmt.Sprintf("%s&signature=%s", params, signature)

	req, err := http.NewRequest("GET", accountInfoURL, nil)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la création de la requête : %v", err)
	}

	req.Header.Set("X-MBX-APIKEY", AppSecret.APIKey)
	req.URL.RawQuery = params

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de l'envoi de la requête : %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("échec de la récupération des informations du compte : %s, réponse : %s", resp.Status, string(bodyBytes))
	}

	var accountInfo AccountInfo
	if err := json.NewDecoder(resp.Body).Decode(&accountInfo); err != nil {
		return nil, fmt.Errorf("erreur lors du décodage de la réponse : %v", err)
	}

	accountInfoCache.info = &accountInfo
	accountInfoCache.lastUpdate = time.Now()

	return &accountInfo, nil
}

func printAccountBalance() error {
	accountInfo, err := getAccountInfo()
	if err != nil {
		return err
	}

	fmt.Println("Soldes du compte :")
	for _, balance := range accountInfo.Balances {
		for _, asset := range AppConfig.Assets {
			if balance.Asset == asset {
				fmt.Printf("Actif : %s, Libre : %s, Bloqué : %s\n", balance.Asset, balance.Free, balance.Locked)
				break
			}
		}
	}
	return nil
}
