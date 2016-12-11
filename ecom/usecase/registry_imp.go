package usecase

import (
	model "github.com/thanhpk/sutu.shop/ecom/model"
)

type Registry struct {
	customerMgt model.ICustomerMgt
}

func (this Registry) Registry(phone string, password string) string {
	customer := &model.Customer{phone: phone, password: password}
	customerid, err := this.customerMgt.Create(customer)
	if err != nil {
		panic (err)
	}

	return customerid
}

func (this Registry) IsDuplicated(phone string) bool {
	c, err := this.customerMgt.MatchByPhone(phone)
	return err == nil
}
