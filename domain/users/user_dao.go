package users

import (
	"fmt"

	"github.com/msd79/bookstore_users-api/datasources/mysql/users_db"
	"github.com/msd79/bookstore_users-api/utils/date_util"
	"github.com/msd79/bookstore_users-api/utils/errors"
	"github.com/msd79/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser                = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	queryFindByStatus           = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?"
)

var (
	userDB = make(map[int64]*User)
)

//Get gets user from the persistent layer
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}

	return nil
}

//Save saves users to the persitent layer
func (user *User) Save() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = date_util.GetNowString()
	insertResult, InsertErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if InsertErr != nil {
		return mysql_utils.ParseError(InsertErr)
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error when tryying to save users %s", err.Error()))
	}
	user.ID = userID

	//user.DateCreated = date_util.GetNowString()

	return nil
}
