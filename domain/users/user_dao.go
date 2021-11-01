package users

import (
	"time"

	"github.com/bookstore_users-api/datasources/mysql/users_db"
	"github.com/bookstore_users-api/utils/errors"
	"github.com/bookstore_users-api/utils/mysql_utils"
)

func (u *User) Save() *errors.RestError {
	stmt, err := users_db.ClientDB.Prepare("INSERT INTO users(first_name,last_name,email,date_created)VALUES (?,?,?,?);")
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()
	insertResult, err := stmt.Exec(u.FirstName, u.LastName, u.Email, time.Now().UTC().String())
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	u.Id = userID
	return nil
}

func (u *User) Get() *errors.RestError {
	stmt, err := users_db.ClientDB.Prepare("SELECT id,first_name,last_name,email,date_created FROM users Where id = ?;")
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()
	result := stmt.QueryRow(u.Id)
	if err := result.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated); err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (u *User) Update() *errors.RestError {
	stmt, err := users_db.ClientDB.Prepare("UPDATE users SET first_name = ? , last_name = ?, email = ? WHERE id = ?;")
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()
	_, err = stmt.Exec(u.FirstName, u.LastName, u.Email, u.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}
