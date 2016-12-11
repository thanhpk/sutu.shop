package usecase


type AddressPO struct {
	Id string
	Phone string
	Address string
}

type ItemPO struct {
	ProductId string
	Quantity int32
}

type ItemDetailsPO struct {
	ItemPO
	TotalPrice int32
}

type OrderPO struct {
	ShippingAddressId string 
	UserId string
	Items []ItemPO
}

type OrderDetails struct {
	Order
	Items []ItemDetailsPO
	TotalPrice int32
}

type IPlaceOrder interface {
	GetOrderDetails(userid string, order Order) OrderDetailsPO
	SaveAddress(userid string, address Address) string
	GetAddresses(useid string) []AddressPO
	PlaceOrder(userid string, ip string, order Order)
}
