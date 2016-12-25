package db

import (
	"github.com/thanhpk/sutu.shop/ecom/model"
	bson "gopkg.in/mgo.v2/bson"
	mgo "gopkg.in/mgo.v2"
)

type MongoProductRepository struct {
	productCollection *mgo.Collection
}

func NewMongoProductRepository(dbname string, dbSession *mgo.Session, context string) *MongoProductRepository {
	pr := MongoProductRepository{}
	pr.productCollection = dbSession.DB(dbname).C(context + "_product")
	return &pr
}

func (me *MongoProductRepository) Update(producttype *model.ProductType) {
	idobj := bson.ObjectIdHex(producttype.Id)
	typeBson, _ := bson.Marshal(producttype)
	var typeM bson.M
	bson.Unmarshal(typeBson, &typeM)
	delete(map[string]interface{}(typeM), "_id")
	err := me.productCollection.UpdateId(idobj, bson.M{"$set": typeM})
	if err != nil {
		panic(err)
	}
}

func (me *MongoProductRepository) Read(idhex string) *model.Product {
	var idobj = bson.ObjectIdHex(idhex)
	product := model.Product{}
	err := me.productCollection.FindId(idobj).One(&product)
	if err != nil {
		return nil
	}
	product.Id = bson.ObjectId(product.Id).Hex()
	return &product
}


func (me *MongoProductRepository) ListByType(typeid string) []model.Product {
	products := []model.Product{}
	err := me.productCollection.Find(bson.M{"TypeId": bson.ObjectIdHex(typeid)}).All(&products)
	if err != nil {
		panic (err)
	}
	return products
}
