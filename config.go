package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Symbol    string   `json:"symbol"`
	Interval  string   `json:"interval"`
	Quantity  float64  `json:"quantity"`
	ShortEMA  int      `json:"shortEMA"`
	LongEMA   int      `json:"longEMA"`
	SleepTime int      `json:"sleepTime"`
	Assets    []string `json:"assets"`
}

type Secret struct {
	APIKey    string `json:"apiKey"`
	APISecret string `json:"apiSecret"`
}

var AppConfig Config
var AppSecret Secret

func LoadConfig() error {
	if err := loadJSON("config.json", &AppConfig); err != nil {
		return err
	}
	return loadJSON("secret.json", &AppSecret)
}

func loadJSON(filename string, v interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(v)
}