package datasources

import (
	"errors"
	"strings"
)

type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

func NewUser(name, email string) *User {
	return &User{
		Name: name,
		Email: email,
	}
}

var (
	Users = make([]*User, 0)
)

func init() {

	Users = append(Users, NewUser("admin", "root@gmail.com"))
}

func (u *User) Cleanup() error {
	u.Name = strings.Trim(u.Name, " ")
	u.Email = strings.Trim(u.Email, " ")

	if len(u.Name) == 0 {
		return errors.New("name can't be empty")
	}
	if len(u.Email) == 0 {
		return errors.New("email can't be empty")
	}

	return nil
}