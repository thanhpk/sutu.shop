package db

import (
	//"fmt"
	"errors"
	bson "gopkg.in/mgo.v2/bson"
	mgo "gopkg.in/mgo.v2"
	"github.com/thanhpk/sutu.shop/ecom/model"
)

func AttachId(obj interface{}, id bson.ObjectId) bson.M{
	doc := bson.D{}
	marshalledObj, _ := bson.Marshal(obj)
	bson.Unmarshal(marshalledObj, &doc)
	docMap := doc.Map()
	docMap["_id"] = id;
	delete(docMap, "Id")
	return docMap
}

type list []interface{}

type IdCustomer struct {
	*model.Customer
	_id bson.ObjectId
}

type MongoCustomerRepository struct {
	customerCollection *mgo.Collection
}
//session.DB(database).C(collection)

func NewMongoCustomerRepository(dbname string, dbSession *mgo.Session, context string) *MongoCustomerRepository {
	cr := MongoCustomerRepository{}
	cr.customerCollection = dbSession.DB(dbname).C(context + "_customer")
	return &cr
}

func (cr MongoCustomerRepository) Create(customer *model.Customer) string {
	newid := bson.NewObjectId()
	err := cr.customerCollection.Insert(AttachId(customer, newid))
	if err != nil {
		panic (err)
	}
	return string(newid)
}

func (cr MongoCustomerRepository) Count(keyword string) int {
	n, err := cr.customerCollection.Find(bson.M{"$or": list{bson.M{"Phone": bson.M{"$regex": keyword}}, bson.M{"Name": bson.M{"$regex": keyword}}}}).Count()

	if err != nil {
		panic(err)
	}
	return n
}
	
func (cr MongoCustomerRepository) List(keyword string, n int, skip int) []model.Customer {
	customers := []model.Customer{}
	err := cr.customerCollection.Find(bson.M{"$or": list{ bson.M{"Phone": bson.M{"$regex": keyword}}, bson.M{"Name": bson.M{"$regex": keyword}}}}).Skip(skip).Limit(n).All(&customers)
	if err != nil {
		panic (err)
	}
	return customers
}

func (cr MongoCustomerRepository) Update(customer *model.Customer) error {
	err := cr.customerCollection.UpdateId(customer.Id, bson.M{"$set": customer})
	if err != nil {
		if err == mgo.ErrNotFound {
			return errors.New("not found")
		}
		panic(err)
	}
	return nil
}

func (cr MongoCustomerRepository) Read(id string) *model.Customer {
	customer := model.Customer{}
	err := cr.customerCollection.FindId(id).One(&customer)
	if err != nil {
		return nil
	}
	return &customer
}

func (cr MongoCustomerRepository) MatchByPhone(phone string) *model.Customer {
	customer := model.Customer{}
	err := cr.customerCollection.Find(bson.M{"Phone": phone}).One(&customer)
	if err != nil {
		return nil
	}
	return &customer
}

func (cr MongoCustomerRepository) MatchByFbUserId(phone string) *model.Customer {
	customer := model.Customer{}
	err := cr.customerCollection.Find(bson.M{"FbUserId": phone}).One(&customer)
	if err != nil {
		return nil
	}
	return &customer
}
