package services

import (
	"github.com/Rahul987725800/golang-microservices/mvc/domain"
)

func GetUser(userID int64) (*domain.User, error) {
	return domain.GetUser(userID)
}