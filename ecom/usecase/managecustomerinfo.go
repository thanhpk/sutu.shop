package usecase

type Customer struct {
	Id string
	Name string
	Email string
	Phone string
	Point int32
	Username string
}

type IManageInfo interface {
	GetInfo(id string) *Customer
}
