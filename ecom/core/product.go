package product


import (
	bson "gopkg.in/mgo.v2/bson"
	mgo "gopkg.in/mgo.v2"
	CA "github.com/thanhpk/sutu.shop/ecom/common/auth"
)

type ProductMgr struct {
	scope string
	database mgo.Database
	collectionName string
}

func NewProductMgr(scope string, database mgo.Database) *ProductMgr {
	productMgr := &ProductMgr{
		scope: scope,
		databae: database
		collectionName
}
