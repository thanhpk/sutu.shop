package usecase

import (
	"time"
)

type AddressL struct {
	Phone string
	Address string
}

type CustomerL struct {
	Id string
	Name string
		
	Point int32
	IsAdmin bool

	Email string
	Phone string
	Addresses []AddressL
}

type ILogin interface {
	AuthByPhone(phone string, password string) (*CustomerL, error)
	AuthByFacebook(userid string, accessToken string) (*CustomerL, error)
}
