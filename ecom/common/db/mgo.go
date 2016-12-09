package mgo

import (
//	mgo "gopkg.in/mgo.v2"
)

type M map[string]interface{}

type D []DocElem

// See the D type.
type DocElem struct {
	Name  string
	Value interface{}
}

type Query interface {
}

type Collection interface {
	FindId(id interface{}) *Query
	DropIndex(key ...string) error
	EnsureIndexKey(key ...string) error 
}

type Database interface {
	C(name string) *Collection
}

type Session interface {
	DB(name string) *Database
}

