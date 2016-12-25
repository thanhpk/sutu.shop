package db

import (
	"github.com/thanhpk/sutu.shop/ecom/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoBrandRepository struct {
	brandCollection *mgo.Collection
}

func NewBrandRepository(dbname string, dbSession *mgo.Session, context string) *MongoBrandRepository {
	br := MongoBrandRepository{}
	br.brandCollection = dbSession.DB(dbname).C(context + "_brand")
	return &br
}

func (me *MongoBrandRepository) Read(id string) *model.Brand {
	idobj := bson.ObjectIdHex(id)
	brand := model.Brand{}
	err := me.brandCollection.FindId(idobj).One(&brand)
	if err != nil {
		return nil
	}
	brand.Id = bson.ObjectId(brand.Id).Hex()
	return &brand
}
