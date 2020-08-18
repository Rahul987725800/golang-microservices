package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Rahul987725800/golang-microservices/mvc/services"
)

// GetUser controller for getting the user
func GetUser(resp http.ResponseWriter, req *http.Request) {
	userID, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		// return bad request to the client
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte("user_id must be a number"))
		return
	}
	user, err := services.GetUser(userID)
	if err != nil {
		// Handle the err and return to the client
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		return
	}
	// return user to client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
