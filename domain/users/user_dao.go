package users

import (
	"fmt"

	"github.com/msd79/bookstore_users-api/datasources/mysql/users_db"
	"github.com/msd79/bookstore_users-api/utils/errors"
	"github.com/msd79/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser                = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	queryFindByStatus           = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?"
)

//Get gets user from the persistent layer
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
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
	insertResult, InsertErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
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

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, updateErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if updateErr != nil {
		return mysql_utils.ParseError(updateErr)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, deleteErr := stmt.Exec(user.ID)
	if deleteErr != nil {
		return mysql_utils.ParseError(deleteErr)
	}
	return nil
}

func (user *User) FindbyStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindByStatus)
	if err != nil {
		if err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
	}
	users, getErr := stmt.Query(status)
	if getErr != nil {
		return nil, mysql_utils.ParseError(getErr)
	}
	defer users.Close()
	result := make([]User, 0)
	for users.Next() {
		var user User
		if err := users.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		result = append(result, user)
	}
	if len(result) == 0 {
		return nil, errors.NewNotFoundError("not matching user found")
	}
	return result, nil
}
