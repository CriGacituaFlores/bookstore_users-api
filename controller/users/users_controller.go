package users

import (
	"github.com/CriGacituaFlores/bookstore_users-api/utils/errors"
	"github.com/CriGacituaFlores/bookstore_users-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/CriGacituaFlores/bookstore_users-api/domain/users"
)

var(counter int)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
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

