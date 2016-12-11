package usecase

import (
	model "github.com/thanhpk/sutu.shop/ecom/model"
)

type Login struct {
	customerMgt model.ICustomerMgt
}

func (this Login) AuthByPhone(phone string, password string) (*CustomerL, error) {
	customerid, err := this.customerMgt.AuthByPhone(phone, password)
	if err != nil {
		return nil, err
	}

	c := this.customerMgt.Read(customerid)
	customer := convertCustomerFromModel(c)
	return customer, nil
}

func convertCustomerFromModel(c model.Customer) *CustomerL {
	customer := new(CustomerL)
	customer.Id = c.Id
	customer.Name = c.Name
	customer.Point = c.Point
	customer.IsAdmin = c.IsAdmin
	customer.Email = c.Email
	customer.Phone = c.Phone
	return customer
}

func (this Login) AuthByFacebook(accesstoken string) (*CustomerL, error) {

	customerid, err := this.customerMgr.AuthByFacebook(accesstoken)
	if err != nil {
		customerid, err := this.customerMgr.CreateFromFacebook(accesstoken)
		if err != nil {
			return nil, err
		}
	}

	c := this.customerMgt.Read(customerid)
	customer := convertCustomerFromModel(c)
	return customer, nil
}


