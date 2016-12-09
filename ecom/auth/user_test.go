package auth

import (
//	"fmt"
	mgo "gopkg.in/mgo.v2"
	"testing"
	imgo "github.com/thanhpk/sutu.shop/ecom/common/db"
)

type QueryMock struct {
}

func (c *CollectionMock) Find(query interface{}) *imgo.Query {
	return nil
}

func (c *CollectionMock) FindId(id interface{}) *imgo.Query {
	return nil
}

type MongoSessionMock struct {
	mgo.Session
}

type DatabaseMock struct {
}

type CollectionMock struct {}

func (*CollectionMock) EnsureIndexKey(key ...string) error {
	return nil
}

func (*CollectionMock) DropIndex(key ...string) error {
	return nil
}

func (DatabaseMock) C(name string) *imgo.Collection {
//	c := imgo.Collection(CollectionMock{})
	//return &c
	return nil
}

func (s *MongoSessionMock) DB(name string) *imgo.Database {
	d := imgo.Database(DatabaseMock{})
	return &d
}

func TestMatchById(t *testing.T) {
//	usermgr := NewUserMgr("111", DatabaseMock{})
	
//	fmt.Println(usermgr)
}

