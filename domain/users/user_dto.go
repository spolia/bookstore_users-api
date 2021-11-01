package users

import "github.com/bookstore_users-api/utils/errors"

type User struct {
	Id          int64  `json: "id"`
	FirstName   string `json: "first_name"`
	LastName    string `json: "last_name"`
	Email       string `json: "email"`
	DateCreated string `json: "date_created"`
}

func (u *User) Validate() *errors.RestError {
	if u.Id == 0 {
		return errors.NewBadRequest("invalid id")
	}
	return nil
}
