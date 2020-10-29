package service

type ISession interface {
	Get(key string)
	Save(key string, data interface{})
}
