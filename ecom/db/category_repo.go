package db

import (
	"github.com/thanhpk/sutu.shop/ecom/model"
	bson "gopkg.in/mgo.v2/bson"
	mgo "gopkg.in/mgo.v2"
)

type MongoCategoryRepository struct {
	categoryCollection *mgo.Collection	
}

func NewMongoCategoryRepository(dbname string, dbSession *mgo.Session, context string) *MongoCategoryRepository {
	cr := MongoCategoryRepository{}
	cr.categoryCollection = dbSession.DB(dbname).C(context + "_category")
	return &cr
}

func (me *MongoCategoryRepository) Read(idhex string) *model.Category {
	idobj := bson.ObjectIdHex(idhex)
	category := model.Category{}
	err := me.categoryCollection.FindId(idobj).One(&category)
	if err != nil {
		return nil
	}
	category.Id = bson.ObjectId(category.Id).Hex()
	return &category
}

func (me *MongoCategoryRepository) All() []model.Category {
	categories := []model.Category{}
	err := me.categoryCollection.Find(nil).All(&categories)
	if err != nil {
		panic (err)
	}
	return categories	
}
