package usecase

type AddressL struct {
	Phone string
	Address string
}

type CustomerL struct {
	Id string
	Name string
		
	Point int
	IsAdmin bool

	Email string
	Phone string
	Addresses []AddressL
}

// usecase dont return error, for example, AuthByPhone alway return customer or panic
type ILogin interface {
	AuthByPhone(phone string, password string) *CustomerL
	AuthByFacebook(accessToken string) *CustomerL
}
