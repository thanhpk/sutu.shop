package usecase

import (
	model "github.com/thanhpk/sutu.shop/ecom/model"
)

type Registry struct {
	customerMgt model.ICustomerMgt
}

func (this Registry) Registry(phone string, password string) string {
	customer := &model.Customer{Phone: phone}
	customerid := this.customerMgt.Create(password, customer)
	return customerid
}

func (this Registry) IsDuplicated(phone string) bool {
	c := this.customerMgt.MatchByPhone(phone)
	return c != nil
}
