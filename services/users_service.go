package services

import (
	"github.com/msd79/bookstore_users-api/domain/users"
	"github.com/msd79/bookstore_users-api/utils/errors"
)

// CreateUser create a user in the database
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
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
