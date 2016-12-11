package model

import (
	bson "gopkg.in/mgo.v2/bson"
	mgo "gopkg.in/mgo.v2"
)

type list []interface{}

type MongoCustomerRepository struct {
	strUserCollection string
	customerCollection mgo.Collection
}

func (cr MongoCustomerRepository) Create(customer *Customer) string {
	customer.Id = string(bson.NewObjectId())
	err := cr.customerCollection.Insert(customer)
	if err != nil {
		panic (err)
	}
	return string(customer.Id)
}

func (cr MongoCustomerRepository) Count(keyword string) int {
	n, err := cr.customerCollection.Find(bson.M{"$or": list{ bson.M{"Phone": bson.M{"$regex": keyword}}, bson.M{"Name": bson.M{"$regex": keyword}}}}).Count()

	if err != nil {
		panic(err)
	}
	return n
}
	
func (cr MongoCustomerRepository) List(keyword string, n int, skip int) []Customer {
	customers := []Customer{}
	err := cr.customerCollection.Find(bson.M{"$or": list{ bson.M{"Phone": bson.M{"$regex": keyword}}, bson.M{"Name": bson.M{"$regex": keyword}}}}).Skip(skip).Limit(n).All(&customers)
	if err != nil {
		panic (err)
	}
	return customers
}

func (cr MongoCustomerRepository) Update(customer *Customer) {
	err := cr.customerCollection.UpdateId(customer.Id, bson.M{"$set": customer})
	if err != nil {
		panic (err)
	}
}

func (cr MongoCustomerRepository) Read(id string) *Customer {
	customer := Customer{}
	err := cr.customerCollection.FindId(id).One(&customer)
	if err != nil {
		return nil
	}
	return &customer
}

func (cr MongoCustomerRepository) MatchByPhone(phone string) *Customer {
	customer := Customer{}
	err := cr.customerCollection.Find(bson.M{"Phone": phone}).One(&customer)
	if err != nil {
		return nil
	}
	return &customer
}

func (cr MongoCustomerRepository) MatchByFbUserId(phone string) *Customer {
	customer := Customer{}
	err := cr.customerCollection.Find(bson.M{"FbUserId": phone}).One(&customer)
	if err != nil {
		return nil
	}
	return &customer
}
