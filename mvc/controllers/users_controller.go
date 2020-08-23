package controllers

import (
	"github.com/Rahul987725800/golang-microservices/mvc/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/Rahul987725800/golang-microservices/mvc/services"
)

// GetUser controller for getting the user
func GetUser(resp http.ResponseWriter, req *http.Request) {
	if req.URL.Query().Get("user_id") != "" {
		userID, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
		if err != nil {
			apiErr := &utils.ApplicationError {
				Message: "user_id must be a number",
				StatusCode: http.StatusBadRequest,
				Code: "bad_request",
			}
			jsonValue, _ := json.Marshal(apiErr)
			// return bad request to the client
			resp.WriteHeader(apiErr.StatusCode)
			resp.Write(jsonValue)
			return
		}
		user, apiErr := services.GetUser(userID)
		if apiErr != nil {
			// Handle the err and return to the client
			resp.WriteHeader(apiErr.StatusCode)
			jsonValue, _ := json.Marshal(apiErr)
			resp.Write(jsonValue)
			return
		}
		// return user to client
		jsonValue, _ := json.Marshal(user)
		resp.Write(jsonValue)
		return
	}
	users := services.GetUsers()
	jsonValue, _ := json.Marshal(users)
	resp.Write(jsonValue)
}
