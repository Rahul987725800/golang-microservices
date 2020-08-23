package domain

import (
	"fmt"
	"net/http"

	"github.com/Rahul987725800/golang-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		1: {ID: 1, FirstName: "Rahul", LastName: "Gupta", Email: "guptarahul@gmail.com"},
		2: {ID: 2, FirstName: "Mehak", LastName: "Gupta", Email: "mehakgupta@gmail.com"},
	}
)

func GetUser(userID int64) (*User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user with id %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}

func GetUsers() map[int64]*User {
	return users
}
