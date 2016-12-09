package auth

import (
	bson "gopkg.in/mgo.v2/bson"
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
		collectionName: scope + "_user",
	}
	
	return userMgr
}

func (u *UserMgr) MatchById(id string) *CA.User {
	c := u.database.C(u.collectionName)
	user := CA.User{}
	err := c.FindId(id).One(&user)
	if err != nil {
		return nil
	}
	return &user
}

func (u *UserMgr) MatchByUserName(username string) *CA.User {
	c := u.database.C(u.collectionName)
	user := CA.User{}
	err := c.Find(bson.M{"username": username}).One(&user)
	if err != nil {
		return nil
	}
	return &user
}

func (u *UserMgr) FindByPhone(phone string) []CA.User {
	c := u.database.C(u.collectionName)
	var users []CA.User

	err := c.Find(bson.M{"phone": phone}).All(&users)
	if err != nil {
		panic (err)
	}
	return users
}

func (u *UserMgr) Create(user *CA.User) string {
	_id := bson.NewObjectId()
	user.Id = string(_id)
	err := u.database.C(u.collectionName).Insert(user)
	if err != nil {
		panic (err)
	}
	return string(_id)
}

func (u *UserMgr) Update(user *CA.User) {
	data, err := bson.Marshal(user)
	if err != nil {
		panic(err)
	}

	err = u.database.C(u.collectionName).UpdateId(user.Id, bson.M{"$set": data})

	if err != nil {
		panic(err)
	}
}
