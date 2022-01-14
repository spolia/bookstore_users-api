package users

import (
	"strings"
	"time"

	"github.com/bookstore_users-api/datasources/mysql/users_db"
	"github.com/bookstore_users-api/logger"
	"github.com/bookstore_users-api/utils/errors"
	"github.com/bookstore_users-api/utils/mysql_utils"
)

// Save insert user into database
func (u *User) Save() *errors.RestError {
	stmt, err := users_db.ClientDB.Prepare("INSERT INTO users(first_name,last_name,email,date_created)VALUES (?,?,?,?);")
	if err != nil {
		logger.Error("preparing insert user statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()
	insertResult, err := stmt.Exec(u.FirstName, u.LastName, u.Email, time.Now().UTC().String())
	if err != nil {
		logger.Error("trying to insert user", err)
		return errors.NewInternalServerError("database error")
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("trying to get last insert id after insert the user", err)
		return errors.NewInternalServerError("database error")
	}

	u.Id = userID
	return nil
}

// Get returns a user from the database
func (u *User) Get() *errors.RestError {
	stmt, err := users_db.ClientDB.Prepare("SELECT id,first_name,last_name,email,date_created FROM users Where id = ?;")
	if err != nil {
		logger.Error("preparing get user statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()
	result := stmt.QueryRow(u.Id)
	if err := result.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated); err != nil {
		logger.Error("trying to get user by id", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (u *User) FindByEmail() *errors.RestError {
	stmt, err := users_db.ClientDB.Prepare("SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=?;")
	if err != nil {
		logger.Error("preparing get user by email statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()
	result := stmt.QueryRow(u.Email)
	if err := result.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated); err != nil {
		if strings.Contains(err.Error(), mysql_utils.ErrorNotRows) {
			return errors.NewNotFound("invalidcredential")
		}

		logger.Error("trying to get user by email", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

// Update updates user data into database
func (u *User) Update() *errors.RestError {
	stmt, err := users_db.ClientDB.Prepare("UPDATE users SET first_name = ? , last_name = ?, email = ? WHERE id = ?;")
	if err != nil {
		logger.Error("preparing update user statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()
	_, err = stmt.Exec(u.FirstName, u.LastName, u.Email, u.Id)
	if err != nil {
		logger.Error("trying to updates the user data", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

// Delete removes a user from the database
func (u *User) Delete() *errors.RestError {
	stmt, err := users_db.ClientDB.Prepare("DELETE FROM users WHERE id = ?;")
	if err != nil {
		logger.Error("preparing delete user statement", err)
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()
	if _, err := stmt.Exec(u.Id); err != nil {
		logger.Error("trying to delete the user by id", err)
		return errors.NewInternalServerError("database error")
	}

	return nil
}
