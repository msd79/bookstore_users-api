package services

import (
	"GoMicroServices/source-code/bookstore_users-api/utils/crypto_utils"

	"github.com/msd79/bookstore_users-api/domain/users"
	"github.com/msd79/bookstore_users-api/utils/date_util"
	"github.com/msd79/bookstore_users-api/utils/errors"
)

// CreateUser create a user in the database
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password)
	user.DateCreated = date_util.GetNowDBFormat()
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

//GetUser gets a user from the persitent layer
func GetUser(userID int64) (*users.User, *errors.RestErr) {
	u := &users.User{ID: userID}

	if err := u.Get(); err != nil {
		return nil, err
	}

	return u, nil
}

//UpdateUser update a user
func UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}
	current.FirstName = user.FirstName
	current.LastName = user.LastName
	current.Email = user.Email

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

// DeleteUser invoke deletion of a user
func DeleteUser(userID int64) *errors.RestErr {
	user := &users.User{ID: userID}
	return user.Delete()
}

// FindbyStatus find and returns list of users with a given status
func Search(status string) ([]users.User, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindbyStatus(status)
}
