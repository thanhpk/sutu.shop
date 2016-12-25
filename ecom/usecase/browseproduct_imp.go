package usecase

import (
	model "github.com/thanhpk/sutu.shop/ecom/model"
	"time"
)

type BrowseProduct struct {
	ProductTypeMgt model.IProductTypeMgt
	BrandMgt model.IBrandMgt
	CategoryMgt model.ICategoryMgt
	VarianceTypeMgt model.IVarianceTypeMgt
	SaleMgt model.ISaleMgt
}

func ConvertSaleFromModel(model *model.Sale) *SaleBP {
	if model == nil {
		return nil
	}
	
	sale := SaleBP{}
	sale.Id = model.Id
	sale.Name = model.Name
	sale.Code = model.Code
	sale.StartTime = model.StartTime
	sale.EndTime = model.EndTime
	sale.CoverImage = model.CoverImage
	sale.QuanlificationCode = model.QuanlificationCode
	return &sale
}

func ConvertCategoryFromModel(model *model.Category) *CategoryBP {
	if model == nil {
		return nil
	}
	
	categorybp := CategoryBP{}
	categorybp.Id = model.Id
	categorybp.Name = model.Name
	categorybp.Url = model.Url
	categorybp.ParentId = model.ParentId

	categorybp.Children = make([]CategoryBP, len(model.Children))
	for idx, child := range model.Children {
		categorybp.Children[idx] = ConvertCategoryFromModel(child)
	}
	
	return &categorybp
}

func ConvertBrandFromModel(model *model.Brand) *BrandBP {
	if model == nil {
		return nil
	}
	
	brandbp := BrandBP{}
	brandbp.Id = model.Id
	brandbp.Name = model.Name
	brandbp.Url = model.Url
	brandbp.LogoImage = model.LogoImage
	return &brandbp
}

func (bp *BrowseProduct) ConvertVariancesFromModel(models []model.Variance) []VarianceBP {
	variancebps := make([]VarianceBP, len(models))
	for idx, variance := range models {
		variancebp := VarianceBP{}
		variancetype := bp.VarianceTypeMgt.Read(variance.VarianceTypeId)
		variancebp.VarianceType = VarianceTypeBP{}
		variancebp.VarianceType.Id = variancetype.Id
		variancebp.VarianceType.Name = variancetype.Name

		variancebp.Value = variance.Value
		variancebps[idx] = variancebp
	}

	return variancebps
}


func (bp *BrowseProduct) ConvertProductFromModel(model *model.Product) *ProductBP {
	if model == nil {
		return nil
	}
	
	productbp := ProductBP{}
	productbp.Id = model.Id
	productbp.Name = model.Name
	productbp.IsInStock = model.Price > 0
	productbp.Price = model.Price
	productbp.SalePrice = model.SalePrice
	productbp.Images = model.Images
	productbp.Variances = bp.ConvertVariancesFromModel(model.Variances)
	
	return &productbp
}

func ConvertProductTypeFromModel(model *model.ProductType, categoryModel *model.Category, brandModel *model.Brand) *ProductTypeBP {
	if model == nil || categoryModel == nil || brandModel == nil {
		return nil
	}
	
	producttypebp := ProductTypeBP{}
	producttypebp.Id = model.Id
	producttypebp.Name = model.Name
	producttypebp.Description = model.Description
	producttypebp.NumberOfView = model.NumberOfView
	producttypebp.NumberOfLove = model.NumberOfLove
	producttypebp.Price = model.Price
	producttypebp.Category = ConvertCategoryFromModel(categoryModel)
	producttypebp.Brand = ConvertBrandFromModel(brandModel)
	return &producttypebp
}

func (bp *BrowseProduct) ListRecentSales() []SaleBP {
	sales := bp.SaleMgt.ListRunningSales()
	salebps := make([]SaleBP, len(sales))
	for idx, sale := range sales {
		salebps[idx] = ConvertSaleFromModel(&sale)
	}
	return salebps
}

func (bp *BrowseProduct) ListMostLovedProductTypes() []ProductTypeBP {
	producttypes := bp.ProductTypeMgt.ListMostLovedProductTypes(30*time.Day)
	producttypebps := make([]ProductTypeBP, len(producttypes))
	for idx, producttype := range producttypes {
		brand := bp.BrandMgt.Read(producttype.BrandId)
		category := bp.CategoryMgt.Read(producttype.CategoryId)
		producttypebps[idx] = ConvertProductTypeFromModel(&producttype, category, brand)
	}
	return producttypebps
}

func (bp *BrowseProduct) ListNewArrivedProductTypes() []ProductTypeBP {
	producttypes := bp.ProductTypeMgt.ListNewArrivedProductTypes()
	producttypebps := make([]ProductTypeBP, len(producttypes))
	for idx, producttype := range producttypes {
		brand := bp.BrandMgt.Read(producttype.BrandId)
		category := bp.CategoryMgt.Read(producttype.CategoryId)
		producttypebps[idx] = ConvertProductTypeFromModel(&producttype, category, brand)
	}
	return producttypebps
}

func (bp *BrowseProduct) GetCategoryTree() *CategoryBP {
	category := bp.CategoryMgt.GetRoot()
	return ConvertCategoryFromModel(category);
}

func (bp *BrowseProduct) GetProductTypesByCategory(categoryid string) []ProductTypeBP {
	producttypes := bp.ProductTypeMgt.ListByCategory(categorystring)
	producttypebps := make([]ProductTypeBP, len(producttypes))
	for idx, producttype := range producttypes {
		brand := bp.BrandMgt.Read(producttype.BrandId)
		category := bp.CategoryMgt.Read(producttype.CategoryId)
		producttypebps[idx] = ConvertProductTypeFromModel(&producttype, category, brand)
	}
	return producttypebps	
}

func (bp *BrowseProduct) GetProductType(id string) *ProductTypeBP {
	producttype := bp.ProductTypeMgt.Read(id)
	return ConvertProductTypeFromModel(producttype)
}

func (bp *BrowseProduct) GetProductsByType(typeid string) []ProductBP {
	products := bp.ProductType.ListProductsByType(typeid)
	productbps := make([]ProductBP, len(products))
	for idx, product := range products {
		producttypebps[idx] = bp.ConvertProductFromModel(&product)
	}
	return producttypebps		
}

func (bp *BrowseProduct) GetProductDetails(productid string) *ProductBP {
	product := bp.ProductTypeMgt.ReadProduct(productid)
	return bp.ConvertProductFromModel(product)
}
