package domain

import (
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(1)
	assert.Nil(t, user, "We were not expecting user with id 0")
	assert.NotNil(t, err, "we were expecting an error when user id is 0")
	if err != nil {
		assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
		assert.EqualValues(t, "user with id 0 was not found", err.Message)
		assert.EqualValues(t, "not_found", err.Code)
	}

}
