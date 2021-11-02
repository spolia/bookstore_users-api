package user

import (
	"net/http"
	"strconv"

	"github.com/bookstore_users-api/domain/users"
	"github.com/bookstore_users-api/services"
	"github.com/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

// Create creates and save a new user
func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, err := services.UserService.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// Get returns a user given a user_id
func Get(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequest("user_id should be a number")
		c.JSON(restErr.Status, restErr)
		return
	}

	user, createError := services.UserService.GetUser(userID)
	if createError != nil {
		c.JSON(createError.Status, createError)
		return
	}

	c.JSON(http.StatusOK, user)
}

// Update updates a user
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

	result, updateErr := services.UserService.UpdateUser(isPartial, user)
	if updateErr != nil {
		c.JSON(updateErr.Status, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// Delete deletes a user given a user_id
func Delete(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequest("user_id should be a number")
		c.JSON(restErr.Status, restErr)
		return
	}
	if err := services.UserService.DeleteUser(userID); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
