package ecom

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"github.com/thanhpk/sutu.shop/common/auth"
)

type ProductType struct {
	id string
	name string
	description string
	no_view int32
	no_love int32
	price int32
	brand_id bson.ObjectId
}

type Brand struct {
	id string
	name string
	description string
	logo_image string
	cover_image string
}

type VarianceType struct {
	id string
	name string
}

type Variance struct {
	id string
	variance_id bson.ObjectId
	value string
}

type Product struct {
	id string

	typy_id string

	price int32
	sale_price int32
	description string
	images []string
	variances []Variance
}

const ORDER_PLACED = 0
const ORDER_CONFIRMED = 1
const ORDER_SHIPPING = 2
const ORDER_SUCCESS = 3

type ShippingAddress auth.Address

type Order struct {
	id string
	shipping_address ShippingAddress
	user_ip string
	user_id string
	status int
	products []Product
	quanties []int32
	is_read bool
	is_paid bool

	create_time time.Time
	modified_time time.Time
}

type Sale struct {
	id string
	start_time time.Time
	end_time time.Time

	cover_image string
	quanlification_code string
}

type Category struct {
	id string
	
	name string
	path string
}
	
