package main

import (
	"sync"
	"time"
)

type Position struct {
	Symbol        string
	EntryPrice    float64
	Quantity      float64
	EntryTime     time.Time
	ShortEMA      float64
	LongEMA       float64
}

type PositionTracker struct {
	mu        sync.RWMutex
	positions map[string]*Position // clé: symbol
}

var GlobalPositionTracker = &PositionTracker{
	positions: make(map[string]*Position),
}

func (pt *PositionTracker) OpenPosition(symbol string, price, quantity, shortEMA, longEMA float64) {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	
	pt.positions[symbol] = &Position{
		Symbol:     symbol,
		EntryPrice: price,
		Quantity:   quantity,
		EntryTime:  time.Now(),
		ShortEMA:   shortEMA,
		LongEMA:    longEMA,
	}
}

func (pt *PositionTracker) ClosePosition(symbol string) *Position {
	pt.mu.Lock()
	defer pt.mu.Unlock()
	
	position := pt.positions[symbol]
	delete(pt.positions, symbol)
	return position
}

func (pt *PositionTracker) HasOpenPosition(symbol string) bool {
	pt.mu.RLock()
	defer pt.mu.RUnlock()
	
	_, exists := pt.positions[symbol]
	return exists
}

func (pt *PositionTracker) GetPosition(symbol string) *Position {
	pt.mu.RLock()
	defer pt.mu.RUnlock()
	
	return pt.positions[symbol]
} 