package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/msd79/bookstore_users-api/utils/errors"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	errorNoRows      = "no rows in result set"
)

//ParseError parses mysql errors and return an appropirate response
func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("No matching record found for the ID")
		}
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewInternalServerError(sqlErr.Message)
	case 1292:
		return errors.NewInternalServerError(sqlErr.Message)
	case 1364:
		return errors.NewInternalServerError(sqlErr.Message)
	}

	return errors.NewInternalServerError("Error proccessing the request")
}
