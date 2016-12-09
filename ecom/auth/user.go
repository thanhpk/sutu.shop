package auth

import (
	imgo "github.com/thanhpk/sutu.shop/ecom/common/db"
	commom_auth "github.com/thanhpk/sutu.shop/ecom/common/auth"
)

type UserMgr struct {
	scope string
	database imgo.Database
	collection string
}

func NewUserMgr(scope string, database imgo.Database) *UserMgr {
	userMgr := &UserMgr{
		scope: scope,
		database: database,
		collection: scope + "_user"	}
	
	return userMgr
}

func (u *UserMgr) MatchById(id string)  commom_auth.User {
	
}
