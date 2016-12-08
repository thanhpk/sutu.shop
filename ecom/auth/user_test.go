package auth

import (
	"fmt"
	mgo "gopkg.in/mgo.v2"
	"testing"
	imgo "github.com/thanhpk/sutu.shop/ecom/common/db"
)

type MongoSessionMock struct {
	mgo.Session
}

func (s MongoSessionMock) DB(name string) *mgo.Database {
	return imgo.Database{}
}

func TestMatchById(t *testing.T) {
	usermgr := NewUserMgr("111", MongoSessionMock{})
	fmt.Println(usermgr)
}
