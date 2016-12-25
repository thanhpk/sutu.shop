package model

import (
	"time"
)

type ProductType struct {
	Id string
	Name string
	Description string
	NumberOfView int
	NumberOfLove int
	Price int
	BrandId string
	CategoryId string
	NewArrived bool
}

type IProductTypeMgt interface {
	Read(typeid string) *ProductType
	ListNewArrivedProductTypes() []ProductType
	ListMostLovedProductTypes(fromtime int) []ProductType
	ReadProduct(productid string) *Product
	ListProductByType(typeid string) []Product
}

type Brand struct {
	Id string `json:"id" bson:"_id,omitempty"`
	Name string
	Url string
	Description string
	LogoImage string
	CoverImage string
}

type IBrandMgt interface {
	Read(id string) *Brand
}

type VarianceType struct {
	Id string `bson:"_id"`
	Name string
}

type IVarianceTypeMgt interface {
	Read(id string) *VarianceType
}

type Variance struct {
	VarianceTypeId string
	Value string
}

type Product struct {
	Id string `bson:"_id"`
	Quantity int
	TypeId string
	Name string
	Price int
	SalePrice int
	Description string
	Images []string
	Variances []Variance
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
	Id string `bson:"_id,omitempty"`
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

type ISaleMgt interface {
	ListRunningSales() []Sale
}

type Sale struct {
	Id string `bson:"_id"`
	Name string
	Code string
	StartTime time.Time
	EndTime time.Time
	CoverImage string
	QuanlificationCode string
}

type Category struct {
	Id string `bson:"_id"`
	Name string
	Url string
	Children []*Category
	ParentId string
}

type ICategoryMgt interface {
	Read(id string) *Category
	GetRoot() *Category
}

type Address struct {
	Phone string
	Address string
}

type Customer struct {
	Id string `bson:"_id"`
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
