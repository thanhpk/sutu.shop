package usecase

type AddressPO struct {
	Id string
	Phone string
	Address string
}

type ItemPO struct {
	ProductId string
	Quantity int
}

type ItemDetailsPO struct {
	ItemPO
	TotalPrice int
}

type OrderPO struct {
	ShippingAddressId string 
	UserId string
	Items []ItemPO
}

type OrderDetailsPO struct {
	OrderPO
	Items []ItemDetailsPO
	TotalPrice int
}

type IPlaceOrder interface {
	GetOrderDetails(userid string, order OrderPO) OrderDetailsPO
	SaveAddress(userid string, address AddressPO) string
	GetAddresses(useid string) []AddressPO
	PlaceOrder(userid string, ip string, order OrderPO)
}
