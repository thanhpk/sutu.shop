package auth

type Address struct {
	Id string
	Phone string
	Address string
}

type User struct {
	Id string
	Name string
	Email string
	Phone string
	HashedPassword string
	Point int32
	IsAdmin bool
	Username string

	Addresses []Address
}

type Auth interface {
	Authenticate(id string, password string) bool
	CanAccess(user_id string, action string) bool
}

type UserDb interface {
	MatchById(id string) User
	MatchByUsername(username string) User
	MatchByPhone(phone string) User

	Create(user User) string // return user id
	Update(user User) // id never change
	Delete(id string)
}

