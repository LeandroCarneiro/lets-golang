package models

import "time"

type RatesHistory struct {
	Id           int
	CurrencyFrom string
	CurrencyTo   string
	Rate         float32
	Date         time.Time
}
