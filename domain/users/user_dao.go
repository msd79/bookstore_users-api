package users

import (
	"fmt"

	"github.com/msd79/bookstore_users-api/logger"

	"github.com/msd79/bookstore_users-api/datasources/mysql/users_db"
	"github.com/msd79/bookstore_users-api/utils/errors"
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
		logger.Error("Error when preparing get user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("Error when trying to get user by ID", getErr)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

//Save saves users to the persitent layer
func (user *User) Save() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("Error when trying prepare save user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	insertResult, InsertErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if InsertErr != nil {
		logger.Error("Error when trying insert user in to", InsertErr)
		return errors.NewInternalServerError("database error")
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("Error when trying get last insert ID", err)
		return errors.NewInternalServerError("database error")
	}
	user.ID = userID

	//user.DateCreated = date_util.GetNowString()

	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("Error when  prepare update user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	_, updateErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if updateErr != nil {
		logger.Error("Error when trying to update user", updateErr)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("Error when  prepare delete user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	_, deleteErr := stmt.Exec(user.ID)
	if deleteErr != nil {
		logger.Error("Error when  trying to delete user ", deleteErr)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) FindbyStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindByStatus)
	if err != nil {
		logger.Error("Error when  prepareing find user statement", err)
		return nil, errors.NewInternalServerError("database error")
	}
	users, getErr := stmt.Query(status)
	if getErr != nil {
		logger.Error("Error when finding user", getErr)
		return nil, errors.NewInternalServerError("database error")
	}
	defer users.Close()
	result := make([]User, 0)
	for users.Next() {
		var user User
		if err := users.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("Error when trying to scan user row into user struct", err)
			return nil, errors.NewInternalServerError("database error")
		}
		result = append(result, user)
	}
	if len(result) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("not matching user found with status %s", status))
	}
	return result, nil
}
