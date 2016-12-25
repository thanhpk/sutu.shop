package usecase

import (
	"time"
)

type SaleBP struct {
	Id string
	Name string
	Code string
	StartTime time.Time
	EndTime time.Time
	CoverImage string
	QuanlificationCode string
}

type VarianceTypeBP struct {
	Id string
	Name string
}

type VarianceBP struct {
	VarianceType VarianceTypeBP
	Value string
}

type BrandBP struct {
	Id string
	Name string
	Url string
	LogoImage string
}

type ProductTypeBP struct {
	Id string
	Name string
	Description string
	NumberOfView int
	NumberOfLove int
	Price int
	Category *CategoryBP
	Brand *BrandBP
}

type ProductBP struct {
	Id string
	IsInStock bool
	Name string
	Price int
	SalePrice int
	Images []string
	Variances []VarianceBP
}

type ProductDetailsBP struct {
	ProductBP
	Description string
}

type CategoryBP struct {
	Id string
	
	Name string
	Url string
	Parent *CategoryBP
}

type IBrowseProduct interface {
	ListRecentSales() []SaleBP
	ListMostLovedProductTypes() []ProductTypeBP
	ListNewArrivedProductTypes() []ProductTypeBP

	GetCategoryTree() *CategoryBP
	GetProductTypesByCategory(categoryid string) []ProductTypeBP
	GetProductType(id string) *ProductTypeBP
	GetProductsByType(typeid string) []ProductBP
	GetProductDetails(productid string) *ProductDetailsBP	
}
