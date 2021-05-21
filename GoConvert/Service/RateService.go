package Service

import (
	d "GoConvert/Domain"

	"github.com/golobby/container/v2/pkg/container"
)

var _repo IRepository

func init() {
	container.New().Bind(&_repo)
}

func Find(id int) d.RatesHistory {
	var result d.RatesHistory
	_repo.Get(id, &result)

	return result
}

func Add(data d.RatesHistory) error {
	err := _repo.Add(&data, "insert")

	return err
}
