package model

type IProductRepository interface {
	Read(id string) *Product
	ListByType(typeid string) []Product
}

type IProductTypeRepository interface {	
//	Create(*ProductType) string
	Update(*ProductType)
	Read(id string) *ProductType
	ListByArrived() []ProductType
	ListByLove() []ProductType
	ListByCategory(category string) []ProductType
}

type ProductTypeMgt struct {
	Repo IProductTypeRepository
	ProductRepo IProductRepository
}

func (me *ProductTypeMgt) Love(typeid string) {
	producttype := me.Read(typeid)
	if producttype == nil {
		panic ("not found product type " + typeid)
	}
	producttype.NumberOfLove = producttype.NumberOfLove + 1
	me.Repo.Update(producttype)
}

func (me *ProductTypeMgt) View(typeid string) {
	producttype := me.Read(typeid)
	if producttype == nil {
		panic ("not found product type " + typeid)
	}
	producttype.NumberOfView = producttype.NumberOfView + 1
	me.Repo.Update(producttype)
}

func (me *ProductTypeMgt) Read(typeid string) *ProductType {
	return me.Repo.Read(typeid)
}

func (me *ProductTypeMgt) ListNewArrivedProductTypes() []ProductType {
	return me.Repo.ListByArrived()
}

func (me *ProductTypeMgt) ListMostLovedProductTypes(fromtime int) []ProductType {
	return me.Repo.ListByLove()
}

func (me *ProductTypeMgt) ListByCategory(categoryid string) []ProductType {
	return me.Repo.ListByCategory(categoryid)
}

func (me *ProductTypeMgt) ReadProduct(productid string) *Product {
	return me.ProductRepo.Read(productid)
}

func (me *ProductTypeMgt) ListProductsByType(typeid string) []Product {
	return me.ProductRepo.ListByType(typeid)
}
