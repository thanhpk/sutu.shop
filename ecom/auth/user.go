package auth

import (
	imgo "github.com/thanhpk/sutu.shop/ecom/common/db"
	commom_auth "github.com/thanhpk/sutu.shop/ecom/common/auth"
)

type UserMgr struct {
	scope string
	session imgo.Session
	database string
}

func NewUserMgr(scope string, session imgo.Session) *UserMgr {
	userMgr := &UserMgr{
		scope: scope,
		session: session,
		database: scope + "_user"	}
	

	return userMgr
}

func (u *UserMgr) MatchById(id string)  commom_auth.User {
	panic ("not implemented")
}
