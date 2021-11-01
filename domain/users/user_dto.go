package users

import (
	"strings"

	"github.com/bookstore_users-api/utils/errors"
)

type User struct {
	Id          int64  `json: "id"`
	FirstName   string `json: "first_name"`
	LastName    string `json: "last_name"`
	Email       string `json: "email"`
	DateCreated string `json: "date_created"`
}

func (u *User) Validate() *errors.RestError {
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)

	u.Email = strings.TrimSpace(u.Email)
	if u.Email == "" {
		return errors.NewBadRequest("invalid email address")
	}

	return nil
}
