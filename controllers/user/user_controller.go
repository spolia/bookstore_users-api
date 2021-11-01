package user

import (
	"net/http"
	"strconv"

	"github.com/bookstore_users-api/domain/users"
	"github.com/bookstore_users-api/services"
	"github.com/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, err := services.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequest("user_id should be a number")
		c.JSON(restErr.Status, restErr)
		return
	}

	user, createError := services.GetUser(userID)
	if createError != nil {
		c.JSON(createError.Status, createError)
		return
	}

	c.JSON(http.StatusOK, user)
}

func Update(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequest("user_id should be a number")
		c.JSON(restErr.Status, restErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userID
	isPartial := c.Request.Method == http.MethodPatch

	result, updateErr := services.UpdateUser(isPartial, user)
	if updateErr != nil {
		c.JSON(updateErr.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {

}
