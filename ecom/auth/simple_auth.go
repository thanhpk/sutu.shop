package auth

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
	ca "github.com/thanhpk/sutu.shop/ecom/common/auth"
)

var ACTIONS []string = []string{"view_product","edit_product","delete_product","view_order"}

type SimpleAuth struct {
	scope string
	userMgr ca.UserDb
}

func NewSimpleAuth(scope string, u ca.UserDb) SimpleAuth {
	auth := SimpleAuth{scope: scope, userMgr: u}
	return auth
}

func (auth SimpleAuth) Authenticate(id string, password string) bool {
	user := auth.userMgr.MatchById(id)
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	return err == nil
}

func (auth SimpleAuth) HashPasswordAndSalt(password string) string {
	passwordBytes := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword[:])
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if strings.Compare(e, a) == 0 {
			return true
		}
	}
	return false
}

func (auth SimpleAuth) CanAccess(userId string, action string) bool {
	user := auth.userMgr.MatchById(userId)
	if user.IsAdmin == true {
		return true
	} else {
		return Contains(ACTIONS, action)
	}
}
