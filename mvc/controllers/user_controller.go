package controllers

import (
	"encoding/json"
	"github.com/dgomezda/golang-microservices/mvc/services"
	"github.com/dgomezda/golang-microservices/mvc/utils"
	"log"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request){
	userIdParam := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	log.Printf("about the process userid %v", userId)
	if err != nil {
		userErr := &utils.ApplicationError{
			Message: "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		userJsonValue, _ := json.Marshal(userErr)
		resp.WriteHeader(userErr.StatusCode)
		resp.Write(userJsonValue)
		return
	}
	user, apiErr:= services.GetUser(userId)
	if apiErr != nil {
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(apiErr.Message))
		return
	}
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}