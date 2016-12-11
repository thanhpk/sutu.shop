package usecase

import (
	model "github.com/thanhpk/sutu.shop/ecom/model"
)

type Address struct {
	Id string
	Phone string
	Address string
}

type Item struct {
	ProductId string
	Quantity int32
}

type ItemDetails struct {
	Item
	TotalPrice int32
}

type Order struct {
	ShippingAddressId string 
	UserId string
	Items []Item
}

type OrderDetails struct {
	Order
	TotalPrice int32
}

type IPlaceOrder interface {
	GetOrderDetails(userid string, order Order) OrderDetails
	SaveAddress(userid string, address Address) string
	GetAddresses(useid string) []Address
	PlaceOrder(userid string, ip string, order Order)
}
