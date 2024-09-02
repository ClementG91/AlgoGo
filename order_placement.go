package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	orderURL = "https://testnet.binance.vision/api/v3/order"
)

func signRequest(params string) string {
	mac := hmac.New(sha256.New, []byte(AppSecret.APISecret))
	mac.Write([]byte(params))
	return hex.EncodeToString(mac.Sum(nil))
}

func placeOrder(symbol, side, orderType string, quantity, price float64) error {
	timestamp := time.Now().Unix() * 1000
	timeInForce := "GTC"
	params := fmt.Sprintf("symbol=%s&side=%s&type=%s&quantity=%f&price=%f&timeInForce=%s&timestamp=%d", symbol, side, orderType, quantity, price, timeInForce, timestamp)
	signature := signRequest(params)
	params = fmt.Sprintf("%s&signature=%s", params, signature)

	fmt.Println("Paramètres de la requête :", params)

	req, err := http.NewRequest("POST", orderURL, nil)
	if err != nil {
		return fmt.Errorf("erreur lors de la création de la requête : %v", err)
	}

	req.Header.Set("X-MBX-APIKEY", AppSecret.APIKey)
	req.URL.RawQuery = params

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("erreur lors de l'envoi de la requête : %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		return fmt.Errorf("échec du placement de l'ordre : %s, réponse : %s", resp.Status, bodyString)
	}

	fmt.Printf("Ordre placé avec succès : %s %f %s à %f\n", side, quantity, symbol, price)
	return nil
}