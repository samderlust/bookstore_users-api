package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/samderlust/bookstore_users-api/utils/errors"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), "no rows") {
			return errors.NewNotFoundError("no record mathcing given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServerError("error processing request")
}
