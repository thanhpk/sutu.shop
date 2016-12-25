package db

import (
	"github.com/thanhpk/sutu.shop/ecom/model"
	"testing"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
	"time"
	"fmt"
)

func ConnectDb() *mgo.Session {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{"172.25.0.2:27017"},
		Timeout:  10 * time.Second,
	}
	
	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		panic (err)
	}

	return session
}

func TestCRUD(t *testing.T)  {
	mgosession := ConnectDb()
	db := NewMongoCustomerRepository("test", mgosession, "customer")

	customer := model.Customer {
		Phone: "12345",
		HashedPassword: "2344535",
	}
	
	id := db.Create(&customer)

	//test read
	readcustomer := db.MatchByPhone(customer.Phone)
	assert.Equal(t, customer.Phone, readcustomer.Phone)
	assert.Equal(t, customer.HashedPassword, readcustomer.HashedPassword)	
	readcustomer = db.Read(id)
	assert.Equal(t, customer.Phone, readcustomer.Phone)	
	assert.Equal(t, customer.HashedPassword, readcustomer.HashedPassword)
	readcustomer = db.Read(readcustomer.Id)

	assert.Equal(t, customer.Phone, readcustomer.Phone)	
	assert.Equal(t, customer.HashedPassword, readcustomer.HashedPassword)

	readcustomer.Phone="5678"
	ii := readcustomer.Id
	db.Update(readcustomer)

	updatedcustomer := db.Read(ii)
	assert.Equal(t, updatedcustomer.Phone, "5678")
	
	fmt.Println("tf")
}
