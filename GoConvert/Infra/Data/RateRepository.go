package Data

import (
	d "GoConvert/Domain"
	"time"
)

type RateRepository struct {
}

func (rate RateRepository) Get(id int, data interface{}) {
	data = d.NewRate(1, "EUR", "BRL", 6.40, time.Now())
}

func (rate RateRepository) Add(data interface{}, query string) error {
	return nil
}
