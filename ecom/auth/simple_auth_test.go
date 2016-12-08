package auth

import ca "github.com/thanhpk/sutu.shop/ecom/common/auth"
import "testing"
//import "fmt"

type MockUserDb struct {
}

func (u MockUserDb) MatchById(id string) ca.User {
	if id == "1" {
		return ca.User{Id: "1", HashedPassword: "$2a$10$gBoSGsE.4ZV/sL5cn8NVFuhou1/DXZWHeYqbCOqAQ/2umnHGbXw1C", IsAdmin: true}
	}

	if id == "2" {
		return ca.User{Id: "2", HashedPassword: "$2a$10$gBoSGsE.4ZV/sL5cn8NVFuhou1/DXZWHeYqbCOqAQ/2umnHGbXw1C"}
	}
	return ca.User{}
	
}

func (u MockUserDb) MatchByUsername(username string) ca.User {
	return ca.User{}
}

func (u MockUserDb) MatchByPhone(phone string) ca.User {
	return ca.User{}
}

func (u MockUserDb) Create(user ca.User) string {
	return ""
}

func (u MockUserDb) Update(user ca.User) {
}

func (u MockUserDb) Delete(id string) {
}


func TestHashPassword(t *testing.T) {
	
	auth := NewSimpleAuth("123", MockUserDb{})
//	pass := auth.HashPassword("123456");
	
	result := auth.Authenticate("1", "123456")

	if !result {
		t.Error("Expected true, got ", result)
	}

	result = auth.Authenticate("2", "1234567")

	if result {
		t.Error("Expected false, got ", result)
	}
}

func TestCanAccess(t *testing.T) {
	auth := NewSimpleAuth("123", MockUserDb{})
	retTrue := auth.CanAccess("1", "delete_user")
	if !retTrue {
		t.Error("Expected true, got ", retTrue)
	}

	retFalse := auth.CanAccess("2", "delete_user")
	if retFalse {
		t.Error("Expected false, got ", retFalse)
	}
}
