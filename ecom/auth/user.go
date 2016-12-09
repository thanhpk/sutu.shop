package auth

import (
	mgo "gopkg.in/mgo.v2"
	CA "github.com/thanhpk/sutu.shop/ecom/common/auth"
)

type UserMgr struct {
	scope string
	database mgo.Database
	collectionName string
}

func NewUserMgr(scope string, database mgo.Database) *UserMgr {
	userMgr := &UserMgr{
		scope: scope,
		database: database,
		collectionName: scope + "_user"	}
	
	return userMgr
}

func (u UserMgr) MatchById(id string) *CA.User {
	c := u.database.C(u.collectionName)
	user := CA.User{}
	err := c.FindId(id).One(&user)
	if err != nil {
		return nil
	}
	return &user
}
