package users

import (
	"fmt"
	"github.com/bookstore_users-api/utils/errors"
)

var usersDB = make(map[int64]*User)

func (u *User) Save() *errors.RestError {
	current := usersDB[u.Id]
	if current != nil {
		if current.Email == u.Email {
			return errors.NewBadRequest(fmt.Sprintf("email %s already exist", u.Email))
		}
		return errors.NewBadRequest(fmt.Sprintf("user %d already exist", u.Id))
	}
	return nil
}

func (u *User) Get() *errors.RestError {
	result := usersDB[u.Id]
	if result == nil {
		return errors.NewNotFound(fmt.Sprintf("user %d not found", u.Id))
	}

	u.Id = result.Id
	u.Email = result.Email
	u.FirstName = result.FirstName
	u.LastName = result.LastName
	u.DateCreated = result.DateCreated
	return nil
}
