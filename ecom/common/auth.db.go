
package auth

import (
	"gopkg.in/mgo.v2/bson"
)

type Address struct {
	_id bson.ObjectId
	phone string
	address string
}

type User struct {
	id bson.ObjectId
	name string
	email string
	phone string
	hashed_password string
	point int32
	isAdmin bool

	addresses []Address
}

type Auth interface {
	Authenticate()
	CanAccess(user_id bson.ObjectId, action string)
}

type UserDb interface {
	MatchById(id bson.ObjectId) User
	MatchByPhone(phone string) User

	Create(user User) string // return user id
	Update(user User) // id never change
	Delete(id bson.ObjectId)
}

