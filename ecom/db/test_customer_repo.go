package db

import (
	"testing"

)

func test_CRUD(t *testing.T) {
	mgosession := ConnectDb(mgohost, mgousername, mgopassword, mgoport)
	db := NewMongoCustomerRepository("test", mgosession
}
