package usecase

type IRegistry interface {
	IsDuplicated(phone string) bool
	Registry(phone string, password string) string
}
