package usecase

type IRegistry interface {
	Registry(phone string, password string) string
}
