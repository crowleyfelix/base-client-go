package services

type remover interface {
	Remove(key string) error
}
