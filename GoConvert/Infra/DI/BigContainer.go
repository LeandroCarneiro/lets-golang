package DI

import (
	repo "GoConvert/Infra/Data"
	serv "GoConvert/Service"

	"github.com/golobby/container"
)

func Register() {
	container.NewContainer().Singleton(func() serv.IRepository {
		return &repo.RateRepository{}
	})
}
