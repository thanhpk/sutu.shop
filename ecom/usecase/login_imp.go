package usecase

import (
	model "github.com/thanhpk/sutu.shop/ecom/model"
)

type Login struct {
	CustomerMgt model.ICustomerMgt
}

func (this Login) AuthByPhone(phone string, password string) (*CustomerL) {
	c, err := this.CustomerMgt.AuthByPhone(phone, password)
	if err != nil {
		return nil
	}

	customer := convertCustomerFromModel(c)
	return customer
}

func convertCustomerFromModel(c *model.Customer) *CustomerL {
	customer := new(CustomerL)
	customer.Id = c.Id
	customer.Name = c.Name
	customer.Point = c.Point
	customer.IsAdmin = c.IsAdmin
	customer.Email = c.Email
	customer.Phone = c.Phone
	return customer
}

func (this Login) AuthByFacebook(accesstoken string) (*CustomerL) {
	c, err := this.CustomerMgt.AuthByFacebook(accesstoken)
	if err != nil {
		customerid := this.CustomerMgt.CreateFromFacebook(accesstoken)
		if err != nil {
			return nil
		}
		c, err = this.CustomerMgt.Read(customerid)
		if err != nil {
			return nil
		}
	}
	
	customer := convertCustomerFromModel(c)
	return customer
}


