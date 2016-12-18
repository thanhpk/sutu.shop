package adapter

import (
	bson "gopkg.in/mgo.v2/bson"
	mgo "gopkg.in/mgo.v2"
)

type MongoProductRepository struct {
	productCollection mgo.Collection
}
