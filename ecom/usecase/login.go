package usecase

type ILogin interface {
	Login(id string, password string)
}
