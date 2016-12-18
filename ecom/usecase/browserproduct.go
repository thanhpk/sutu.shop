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

type ProductTypeBP struct {
	Id string
	Name string
	Description string
	NumverOfView int32
	NumberOfLove int32
	Price int32
	BrandName string
	CategoryId string
}

type ProductBP struct {
	Id string
	IsInStock bool
	Type ProductTypeBP
	Name string
	Price int32
	SalePrice int32
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
	Path string
	Parent *CategoryBP
}

type IBrowseProduct interface {
	ListRecentSale() []SaleBP
	ListMostLovedProducts() []ProductBP
	ListNewArriveProducts() []ProductBP

	GetCategoryTree() []CategoryBP
	ListProductByCategory() []ProductBP
	GetProductDetails(productid string) *ProductDetailsBP
}
