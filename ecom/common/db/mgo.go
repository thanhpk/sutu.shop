package mgo

import (
//	mgo "gopkg.in/mgo.v2"
)

type Collection interface {
	EnsureIndexKey(key ...string) error 
}

type Database interface {
	C(name string) *Collection
}

type Session interface {
	DB(name string) *Database
}

