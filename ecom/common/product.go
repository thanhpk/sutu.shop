package common

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type ProductType struct {
	_id bson.ObjectId
	name string
	description string

	price int32

}

type Brand struct {
	_id bson.ObjectId
	name string
	description string
	logo_image string
	cover_image string
}

type ShippingAddress struct {
	_id bson.ObjectId
	phone string
	address string
}
	

type VarianceType struct {
	_id bson.ObjectId
	name string
}

type Variance struct {
	variance_id string
	value string
}

type Product struct {
	_id string

	type_id string

	price string
	description string
	images []string
	variances []Variance
}

const ORDER_PLACED = 0
const ORDER_SHIPPING = 0

type Order struct {
	_id bson.ObjectId
	shipping_address ShippingAddress
	user_ip string
	user_id bson.ObjectId
	status int
	products []Product
	quanties []int32
	is_paid bool
	create_time time.Time
	modified_time time.Time
}
