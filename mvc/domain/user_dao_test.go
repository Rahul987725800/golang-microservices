package domain

import (
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)
	if user != nil {
		t.Error("We were not expecting user with id 0")
	}
	if err == nil {
		t.Error("we were expecting an error when user id is 0")
	}
	if err != nil && err.StatusCode != http.StatusNotFound {
		t.Error("we were expecting 404 when user is not found")
	}
}
