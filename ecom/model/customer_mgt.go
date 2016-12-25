package model

import (
	"time"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

type ICustomerRepository interface {
	Create(*Customer) string
	Count(keyword string) int
	List(keyword string, n int, skip int) []Customer
	Update(*Customer)
	Read(id string) *Customer
	MatchByPhone(string) *Customer
	MatchByFbUserId(string) *Customer
}

type FbAppInfo struct {
	Id string
	Name string
}

type FbUserInfo struct {
	Id string
	Email string
	Name string
	Phone string
}

type IFacebookGraphApi interface {
	GetApp(accesstoken string) *FbAppInfo
	GetUser(accesstoken string) *FbUserInfo
}

type CustomerMgt struct {
	FbAppId string
	Repo ICustomerRepository
	Fb IFacebookGraphApi
}

func (c *CustomerMgt) AuthByPhone(phone string, password string) (*Customer, error) {
	customer := c.MatchByPhone(phone)
	if customer == nil { return nil, errors.New("phone number doesn't existed") }

	if customer.HashedPassword == "" {
		return nil, errors.New("user are not login using phone, use fb to login instead")
	}
	
	err := bcrypt.CompareHashAndPassword([]byte(customer.HashedPassword), []byte(password))
	if err != nil { return nil, errors.New("wrong password") }
	return customer, nil
}

func (c *CustomerMgt) Read(id string) (*Customer, error) {
	customer := c.Repo.Read(id)
	if customer == nil {
		return nil, errors.New("user " + id + " not found")
	}
	return customer, nil
}

func (c * CustomerMgt) List(keyword string, n int, skip int) []Customer {
	return c.Repo.List(keyword, n, skip)
}
func (c *CustomerMgt) Count(keyword string) int {
	return c.Repo.Count(keyword)
}

func (c *CustomerMgt) getFbUserSecure(accesstoken string) (*FbUserInfo, error) {
	appInfo := c.Fb.GetApp(accesstoken)

	if appInfo == nil {
		return nil, errors.New("wrong access token")
	}
	
	if appInfo.Id != c.FbAppId {
		return nil, errors.New("wrong facebook appid")
	}

	userInfo := c.Fb.GetUser(accesstoken)
	if userInfo == nil {
		return nil, errors.New("wrong access token")
	}

	return userInfo, nil
}

func (c *CustomerMgt) AuthByFacebook(accesstoken string) (*Customer, error) {
	userInfo, err := c.getFbUserSecure(accesstoken)
	if err != nil {
		return nil, err
	}
	customer := c.Repo.MatchByFbUserId(userInfo.Id)
	if customer == nil {
		return nil, errors.New("customer not found")
	}

	return customer, nil
}

func (c *CustomerMgt) CreateFromFacebook(accesstoken string) string {
	userInfo, err := c.getFbUserSecure(accesstoken)
	if err != nil { panic(err) }

	customer := Customer{
		Name: userInfo.Name,
		Email: userInfo.Email,
		FbUserId: userInfo.Id,
		FbAccessToken: accesstoken,
		CreateTime: time.Now(),
		LastLogin: time.Now(),
	}
	
	return c.Repo.Create(&customer)
}

func hashPassword(password string) string {
	passwordBytes := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword[:])
}

func (c *CustomerMgt) Create(password string, customer *Customer) (string) {
	//do not allow duplicate phone
	if customer.Phone != "" {
		if c.MatchByPhone(customer.Phone) != nil {
			panic("phone number is already existed")
		}
	}
	
	customer.HashedPassword = hashPassword(password)
	customer.CreateTime = time.Now()
	customer.LastLogin = time.Now()
	customer.FbUserId = ""
	customer.FbAccessToken = ""
	customer.IsAdmin = false
	customer.Point = 0
	fmt.Println("gret")
	return c.Repo.Create(customer)
}

func (c *CustomerMgt) MatchByPhone(phone string) *Customer {
	return c.Repo.MatchByPhone(phone)
}
