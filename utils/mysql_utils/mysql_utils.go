package mysql_utils

import (
	"strings"

	"github.com/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	ErrorNotRows = "no rows in result set"
)

func ParseError(err error) *errors.RestError {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNotRows) {
			return errors.NewNotFound("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequest("invalid data")
	}

	return errors.NewInternalServerError("processing request")
}
