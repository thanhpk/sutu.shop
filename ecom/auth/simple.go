package auth

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
	bson "gopkg.in/mgo.v2/bson"
	mgo "gopkg.in/mgo.v2"
)

var ACTIONS []string = []string{"view_product","edit_product","delete_product","view_order"}

type SimpleAuth struct {
	scope string
	userMgr UserMgr
}

type MongoEndPointInfo struct {
	host, database, username, password, collection_prefix string
	port int32
}

func NewSimpleAuth(scope string, u UserMgr) SimpleAuth {
	auth := SimpleAuth{scope: scope, userMgr: u}
}

func (auth SimpleAuth) Authenticate(id string, password string) {
	user := auth.userMgr.MatchById(id)
	err != bcrypt.CompareHashAndPassword(user.hashed_password, password)

	if err != nil {
		panic (err)
	}
}

func (auth SimpleAuth) HashPasswordAndSalt(password string) string) {
	password_byte := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(password_byte, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return hashedPassword
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if strings.Compare(e, a) == 0 {
			return true
		}
	}
	return false
}

func (auth SimpleAuth) CanAccess(user_id bson.ObjectId, action string) bool {
	user := auth.userMgr.MatchById(user_id)
	if user.is_admin == true {
		return true
	} else {
		return Contains(ACTIONS, action)
	}
}
