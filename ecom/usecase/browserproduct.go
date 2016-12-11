package usecase

import (
	"time"
)

type Sale struct {
	Id string
	Name string
	Code string
	StartTime time.Time
	EndTime time.Time
	CoverImage string
	QuanlificationCode string
}

type VarianceType struct {
	Id string
	Name string
}

type Variance struct {
	VarianceType VarianceType
	Value string
}

type ProductType struct {
	Id string
	Name string
	Description string
	NumverOfView int32
	NumberOfLove int32
	Price int32
	BrandName string
}

type Product struct {
	Id string
	IsInStock bool
	Type ProductType
	Name string
	Price int32
	SalePrice int32
	Images []string
	Variances []Variance
}

type ProductDetails struct {
	Product
	Description string
}

type Category struct {
	Id string
	
	Name string
	Path string
	Parent *Category
}

type IBrowserProduct interface {
	ListRecentSale() []Sale
	ListMostLovedProducts() []Product
	ListNewArriveProducts() []Product

	GetCategoryTree() []Category
	ListProductByCategory() []Product
	GetProductDetails(productid string) *ProductDetails
}
