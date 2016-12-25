package model

type IProductRepository interface {
	Read(id string) *Product
	ListByType(tpeid string) []Product
}

type IProductTypeRepository interface {	
	Create(*ProductType) string
	Count(search string) string
	List(keyword string, n int, p int) []ProductType
	Update(*ProductType) string
	Read(id string) *ProductType
	ListByArrived() []ProductType
	ListByLove() []ProductType
}

type ProductTypeMgt struct {
	Repo IProductTypeRepository
	ProductRepo IProductRepository
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

func (me *ProductTypeMgt) ReadProduct(productid string) *Product {
	return me.ProductRepo.Read(productid)
}

func (me *ProductTypeMgt) ListProductByType(typeid string) []Product {
	return me.ProductRepo.ListByType(typeid)
}
