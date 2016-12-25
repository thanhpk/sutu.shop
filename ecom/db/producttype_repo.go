package db


import (
	"github.com/thanhpk/sutu.shop/ecom/model"
	bson "gopkg.in/mgo.v2/bson"
	mgo "gopkg.in/mgo.v2"
)

type MongoProductTypeRepository struct {
	typeCollection *mgo.Collection
	
}

func NewMongoProductTypeRepository(dbname string, dbSession *mgo.Session, context string) *MongoProductTypeRepository {
	pr := MongoProductTypeRepository{}
	pr.typeCollection = dbSession.DB(dbname).C(context + "_type")
	return &pr
}

func (me *MongoProductTypeRepository) Update(producttype *model.ProductType) {
	idobj := bson.ObjectIdHex(producttype.Id)
	typeBson, _ := bson.Marshal(producttype)
	var typeM bson.M
	bson.Unmarshal(typeBson, &typeM)
	delete(map[string]interface{}(typeM), "_id")
	
	err := me.typeCollection.UpdateId(idobj, bson.M{"$set": typeM})
	if err != nil {
		panic(err)
	}
}

func (me *MongoProductTypeRepository) Read(idhex string) *model.ProductType {
	var idobj = bson.ObjectIdHex(idhex)
	producttype := model.ProductType{}
	err := me.typeCollection.FindId(idobj).One(&producttype)
	if err != nil {
		return nil
	}
	producttype.Id = bson.ObjectId(producttype.Id).Hex()
	return &producttype
}

func (me *MongoProductTypeRepository) ListByArrived() []model.ProductType {
	producttypes := []model.ProductType{}
	err := me.typeCollection.Find(nil).Sort("-NewArrived").Skip(0).Limit(10).All(&producttypes)
	if err != nil {
		panic (err)
	}
	return producttypes
}

func (me *MongoProductTypeRepository) ListByLove() []model.ProductType {
	producttypes := []model.ProductType{}
	err := me.typeCollection.Find(nil).Sort("-NumberOfLove").Skip(0).Limit(10).All(&producttypes)
	if err != nil {
		panic (err)
	}
	return producttypes
}
