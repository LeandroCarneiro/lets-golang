package Domain

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("init domain")
}

type RatesHistory struct {
	Id           int
	CurrencyFrom string
	CurrencyTo   string
	Rate         float32
	Date         time.Time
}

func NewRate(id int, from string, to string, rate float32, date time.Time) *RatesHistory {
	r := RatesHistory{
		Id:           id,
		CurrencyFrom: from,
		CurrencyTo:   to,
		Rate:         rate,
		Date:         date,
	}

	return &r
}
