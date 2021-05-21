package Service

type IRepository interface {
	Add(data interface{}, query string) error
	Get(id int, data interface{})
}
