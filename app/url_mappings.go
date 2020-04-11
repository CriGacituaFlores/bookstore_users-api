package app

import (
	"github.com/CriGacituaFlores/bookstore_users-api/controller/ping"
	"github.com/CriGacituaFlores/bookstore_users-api/controller/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
}