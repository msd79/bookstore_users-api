package users

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/msd79/bookstore_users-api/domain/users"
	"github.com/msd79/bookstore_users-api/services"
	"github.com/msd79/bookstore_users-api/utils/errors"
)

func GetUserIdParam(userIdParam string) (int64, *errors.RestErr) {
	userID, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("invalid user id")
	}
	return userID, nil
}

//Get returns a user
func Get(c *gin.Context) {
	userID, idErr := GetUserIdParam(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	u, err := services.GetUser(userID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, u)
}

// Create creates a user
func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO: Handle json bindin error here
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {

		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)

}

//Update updates a user
func Update(c *gin.Context) {

	userID, idErr := GetUserIdParam(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO: Handle json bindin error here
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user.ID = userID
	result, err := services.UpdateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

//Delete deletes a user
func Delete(c *gin.Context) {
	userID, idErr := GetUserIdParam(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	err := services.DeleteUser(userID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return

	}
	c.JSON(http.StatusOK, users)
}
