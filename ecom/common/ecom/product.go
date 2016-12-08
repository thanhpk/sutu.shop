package ecom

import (
	"time"
	"github.com/thanhpk/sutu.shop/ecom/common/auth"
)

type ProductType struct {
	Id string
	Name string
	Description string
	NumverOfView int32
	NumberOfLove int32
	Price int32
	BrandId string
}

type Brand struct {
	Id string
	Name string
	Description string
	LogoImage string
	CoverImage string
}

type VarianceType struct {
	Id string
	Name string
}

type Variance struct {
	Id string
	VarianceId string
	Value string
}

type Product struct {
	Id string

	TypeId string

	Price int32
	SalePrice int32
	Description string
	Images []string
	Variances []Variance
}

const ORDER_PLACED = 0
const ORDER_CONFIRMED = 1
const ORDER_SHIPPING = 2
const ORDER_SUCCESS = 3

type ShippingAddress auth.Address

type Order struct {
	Id string
	ShippingAddress ShippingAddress
	UserIp string
	UserIxd string
	Status int
	Products []Product
	Quanties []int32
	IsRead bool
	IsPaid bool

	CreateTime time.Time
	LastModifiedTime time.Time
}

type Sale struct {
	Id string
	StartTime time.Time
	EndTime time.Time

	CoverImage string
	QuanlificationCode string
}

type Category struct {
	Id string
	
	Name string
	Path string
}
	
