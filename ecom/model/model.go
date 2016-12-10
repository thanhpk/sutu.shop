package model

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
	Category Category
}

type IProductTypeRepository interface {
	Create(*ProductType) string
	Count(search string) string
	List(keyword string, n int32, p int32) []ProductType
	Update(*ProductType) string
	Read(id string) *ProductType
}

type Brand struct {
	Id string
	Name string
	Description string
	LogoImage string
	CoverImage string
}

type IBrandRepository interface {
	Create(*Brand) string
	Count(search string) string
	List(keyword string, n int32, p int32) []Brand
	Update(*Brand) string
	Read(id string) *Brand
}

type VarianceType struct {
	Id string
	Name string
}

type IVarianceTypeRepository interface {
	Create(*VarianceType) string
	Count(search string) string
	List(keyword string, n int32, p int32) []VarianceType
	Update(*VarianceType) string
	Read(id string) *VarianceType
}

type Variance struct {
	Id string
	VarianceId string
	Value string
}

type Product struct {
	Id string
	TypeId string
	Name string
	Price int32
	SalePrice int32
	Description string
	Images []string
	Variances []Variance
}

type IProductRepository interface {
	Create(*Product) string
	Count(search string) string
	List(keyword string, n int32, p int32) []Product
	Update(*Product) string
	Read(id string) *Product
}

const ORDER_PLACED = 0
const ORDER_CONFIRMED = 1
const ORDER_SHIPPING = 2
const ORDER_SUCCESS = 3

type Item struct {
	Product Product
	Quantity int32
}

type Order struct {
	Id string
	Code string
	ShippingAddress Address
	UserIp string
	UserId string
	Status int
	Items []Item
	IsRead bool
	IsPaid bool
	CreateTime time.Time
	LastModifiedTime time.Time
}

type IOrderRepository interface {
	Create(*Order) string
	Count(keyword string) string
	List(keyword string, n int32, p int32) []Order
	Update(*Order) string
	Read(id string) *Order
	Match(code string) *Order
}

type Sale struct {
	Id string
	Name string
	Code string
	StartTime time.Time
	EndTime time.Time
	CoverImage string
	QuanlificationCode string
}

type ISaleRepository interface {
	Cerate(*Sale) string
	Count(keyword string) string
	List(keyword string, n int32, p int32) []Sale
	Update(*Sale) string
	Read(id string) *Sale
	Match(code string) *Sale
}

type Category struct {
	Id string
	
	Name string
	Path string
	Parent *Category
}

type ICategoryRepository interface {
	Cerate(*Category) string
	Count(keyword string) string
	List(keyword string, n int32, p int32) []Category
	Update(*Category) string
	Read(id string) *Category
	Match(code string) *Category
}

type Address struct {
	Id string
	Phone string
	Address string
}

type IAddressRepository interface {
	Cerate(*Address) string
	Count(keyword string) string
	List(keyword string, n int32, p int32) []Address
	Update(*Address) string
	Read(id string) *Address
	Match(code string) *Address
}

type Customer struct {
	Id string
	Name string
	Email string
	Phone string
	HashedPassword string
	Point int32
	IsAdmin bool
	Username string

	Addresses []Address
}

type Auth interface {
	Authenticate(id string, password string) bool
	CanAccess(user_id string, action string) bool
}

type ICustomerRepository interface {

	Create(*Customer) string
	Count(keyword string) int32
	List(keyword string, n int32, p int32) []Customer
	Update(*Customer) string
	Read(id string) *Customer
	
	MatchByUsername(string) *Customer	
}


