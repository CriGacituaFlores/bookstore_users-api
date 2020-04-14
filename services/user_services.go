package services

import (
	"github.com/CriGacituaFlores/bookstore_users-api/utils/errors"
	"github.com/CriGacituaFlores/bookstore_users-api/domain/users"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}