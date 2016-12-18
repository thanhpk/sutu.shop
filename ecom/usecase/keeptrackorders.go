package usecase

import (
	"time"
)

type ProductTypeKP struct {
	Id string
	Name string
	Description string
	NumverOfView int32
	NumberOfLove int32
	Price int32
	BrandId string
	CategoryId string
}

type ProductKP struct {
	Id string
	IsInStock bool
	Type ProductTypeKP
	Name string
	Price int32
	Images []string
	Variances []VarianceKP
}

type VarianceTypeKP struct {
	Id string
	Name string
}

type VarianceKP struct {
	VarianceType VarianceTypeKP
	Value string
}

type ItemDetailsKP struct {
	Product ProductKP
	Quantity int32
	TotalPrice int32
}

type AddressKP struct {
	Id string
	Phone string
	Address string
}

type OrderKP struct {
	Id string
	Code string
	ShippingAddress AddressKP
	UserIp string
	Status int
	Items []ItemDetailsKP
	IsPaid bool
	CreateTime time.Time
	LastModifiedTime time.Time
	TotalPrice int32
}

type IKeepTrackOrders interface {
	ListAllOrder() []OrderKP
}
