package services

import (
	"github.com/bookstore_users-api/domain/users"
	"github.com/bookstore_users-api/utils/errors"
)

var UserService userClient = &userservice{}

type userservice struct {
}

type userClient interface {
	GetUser(int64) (*users.User, *errors.RestError)
	CreateUser(users.User) (*users.User, *errors.RestError)
	UpdateUser(bool, users.User) (*users.User, *errors.RestError)
	DeleteUser(int64) *errors.RestError
}

func (s *userservice) GetUser(userID int64) (*users.User, *errors.RestError) {
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

func (s *userservice) CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *userservice) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestError) {
	current, err := s.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if isPartial {
		if current.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if current.LastName != "" {
			current.LastName = user.LastName
		}
		if current.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (s *userservice) DeleteUser(userID int64) *errors.RestError {
	current, err := s.GetUser(userID)
	if err != nil {
		return err
	}

	return current.Delete()
}
