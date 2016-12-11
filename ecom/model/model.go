package model

import (
	"time"
)

type ProductType struct {
	Id string `json:"id" bson:"_id,omitempty"`
	Name string
	Description string
	NumverOfView int
	NumberOfLove int
	Price int
	BrandId string
	CategoryId string
}

type IProductTypeMgt interface {
	
}

type ProductTypeMgt struct {
	
}

type IProductTypeRepository interface {
	Create(*ProductType) string
	Count(search string) string
	List(keyword string, n int, p int) []ProductType
	Update(*ProductType) string
	Read(id string) *ProductType
}

type Brand struct {
	Id string `json:"id" bson:"_id,omitempty"`
	Name string
	Description string
	LogoImage string
	CoverImage string
}

type IBrandRepository interface {
	Create(*Brand) string
	Count(search string) string
	List(keyword string, n int, p int) []Brand
	Update(*Brand) string
	Read(id string) *Brand
}

type VarianceType struct {
	Id string `json:"id" bson:"_id,omitempty"`
	Name string
}

type IVarianceTypeRepository interface {
	Create(*VarianceType) string
	Count(search string) string
	List(keyword string, n int, p int) []VarianceType
	Update(*VarianceType) string
	Read(id string) *VarianceType
}

type Variance struct {
	VarianceTypeId string
	Value string
}

type Product struct {
	Id string `json:"id" bson:"_id,omitempty"`
	Quantity int
	TypeId string
	Name string
	Price int
	SalePrice int
	Description string
	Images []string
	Variances []Variance
}

type IProductRepository interface {
	Create(*Product) string
	Count(search string) string
	List(keyword string, n int, p int) []Product
	Update(*Product) string
	Read(id string) *Product
}

const ORDER_PLACED = 0
const ORDER_CONFIRMED = 1
const ORDER_SHIPPING = 2
const ORDER_SUCCESS = 3

type Item struct {
	ProductId string
	Quantity int
}

type Order struct {
	Id string `json:"id" bson:"_id,omitempty"`
	Code string
	ShippingAddressId string
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
	List(keyword string, n int, p int) []Order
	Update(*Order) string
	Read(id string) *Order
	Match(code string) *Order
}

type Sale struct {
	Id string `json:"id" bson:"_id,omitempty"`
	Name string
	Code string
	StartTime time.Time
	EndTime time.Time
	CoverImage string
	QuanlificationCode string
}

type ISaleRepository interface {
	Create(*Sale) string
	Count(keyword string) string
	List(keyword string, n int, p int) []Sale
	Update(*Sale) string
	Read(id string) *Sale
	Match(code string) *Sale
}

type Category struct {
	Id string `json:"id" bson:"_id,omitempty"`
	
	Name string
	Path string
	Parent *Category
}

type ICategoryRepository interface {
	Create(*Category) string
	Count(keyword string) string
	List(keyword string, n int, p int) []Category
	Update(*Category) string
	Read(id string) *Category
	Match(code string) *Category
}

type Address struct {
	Phone string
	Address string
}

type IAddressRepository interface {
	Create(*Address) string
	Count(keyword string) string
	List(keyword string, n int, p int) []Address
	Update(*Address) string
	Read(id string) *Address
	Match(code string) *Address
}

type Customer struct {
	Id string `json:"id" bson:"_id,omitempty"`
	Name string
	HashedPassword string
	Point int
	IsAdmin bool

	FbUserId string
	FbAccessToken string
	
	Email string
	Phone string
	Addresses []Address

	CreateTime time.Time
	LastLogin time.Time
}


type Auth interface {
	Authenticate(id string, password string) bool
	CanAccess(user_id string, action string) bool
}

type ICustomerRepository interface {
	Create(*Customer) string
	Count(keyword string) int
	List(keyword string, n int, skip int) []Customer
	Update(*Customer) error
	Read(id string) *Customer
	MatchByPhone(string) *Customer
	MatchByFbUserId(string) *Customer
}

type ICustomerMgt interface {
	AuthByPhone(phone string, password string) (*Customer, error)
	Read(id string) (*Customer, error)
	AuthByFacebook(accesstoken string) (*Customer, error)
	CreateFromFacebook(accesstoken string) string //panic
	Create(password string, cus *Customer) string //panic
	MatchByPhone(string) *Customer
	List(keyword string, n int, skip int) []Customer
	Count(keyword string) int
}
