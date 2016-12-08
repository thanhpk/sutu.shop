
package common

type User struct {
	id string
	name string
	email string
	phone string
	hashed_password string

	isAdmin bool
}

type UserDb interface {
	MatchById(id string) User
	MatchByPhone(phone string) User

	Create(user User) string // return user id
	Update(user User) // id never change
	Delete(id string)
}


