package auth

import (
//	bson "gopkg.in/mgo.v2/bson"
	mgo "gopkg.in/mgo.v2"
	commom_auth "github.com/thanhpk/sutu.shop/ecom/common/auth"
)

type UserMgr struct {
	scope string
	session mgo.Session
	database string
}

func NewUserMgr(scope string) *UserMgr {
	userMgr := &UserMgr{}

	return userMgr

}

func (u *UserMgr) MatchById(id string)  commom_auth.User {
	panic ("not implemented")
}
