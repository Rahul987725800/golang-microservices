package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	// Initialization:
	// Execution:
	user, err := GetUser(0)
	// Validation:
	assert.Nil(t, user, "We were not expecting user with id 0")
	assert.NotNil(t, err, "we were expecting an error when user id is 0")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user with id 0 was not found", err.Message)
	assert.EqualValues(t, "not_found", err.Code)
}
func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(1)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.ID)
	assert.EqualValues(t, "Rahul", user.FirstName)
	assert.EqualValues(t, "Gupta", user.LastName)
	assert.EqualValues(t, "guptarahul@gmail.com", user.Email)
}

func TestGetUsers(t *testing.T) {
	allUsers := GetUsers()
	assert.NotNil(t, users)
	assert.EqualValues(t, allUsers, users)
}
