package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var(counter int)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

