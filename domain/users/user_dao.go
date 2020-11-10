package users

import (
	"fmt"

	"github.com/msd79/bookstore_users-api/utils/date_util"
	"github.com/msd79/bookstore_users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

//Get gets user from the persistent layer
func (user *User) Get() *errors.RestErr {
	result := userDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}

	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

//Save saves users to the persitent layer
func (user *User) Save() *errors.RestErr {
	current := userDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError((fmt.Sprintf("User email %s already registered", user.Email)))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User with %d exist", user.ID))
	}

	user.DateCreated = date_util.GetNowString()
	userDB[user.ID] = user
	return nil
}
