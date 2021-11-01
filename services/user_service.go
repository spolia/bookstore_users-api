package services

import (
	"github.com/bookstore_users-api/domain/users"
	"github.com/bookstore_users-api/utils/errors"
)

func GetUser(userID int64) (*users.User, *errors.RestError) {
	if userID <= 0 {
		return nil, errors.NewBadRequest("invalid user_id")
	}

	user := &users.User{
		Id: userID,
	}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}
func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func FindUser() {

}

