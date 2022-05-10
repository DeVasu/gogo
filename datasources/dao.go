package datasources

import (
	"errors"
	"strings"
)

func (u *User) Add() error {

	if err := u.Cleanup(); err != nil {
		return err
	}
	

	for _, ele := range Users {
		if ele.Email == u.Email {
			return errors.New("user with same email exists")
		}
		if ele.Name == u.Name {
			return errors.New("user with same name exists")
		}
	}

	Users = append(Users, u)

	return nil

}
func (u *User) GetByName() error {

	u.Name = strings.Trim(u.Name, " ")
	if len(u.Name) == 0 {
		return errors.New("name in request can't be empty")
	}

	for _, ele := range Users {
		if ele.Name == u.Name {
			u.Email = ele.Email
			return nil
		}
	}

	return errors.New("user does not exist")
}
