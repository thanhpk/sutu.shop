package auth

import (
	"gopkg.in/mgo.v2/bson"
)

type Address struct {
	id string
	phone string
	address string
}

type User struct {
	id string
	name string
	email string
	phone string
	hashed_password string
	point int32
	is_admin bool
	username string

	addresses []Address
}

type Auth interface {
	Authenticate()
	CanAccess(user_id bson.ObjectId, action string)
}

type UserDb interface {
	MatchById(id string) User
	MatchByUsername(username string) User
	MatchByPhone(phone string) User

	Create(user User) string // return user id
	Update(user User) // id never change
	Delete(id bson.ObjectId)
}

