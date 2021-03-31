package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/nwneisen/MicroservicesInGo/Section2-MVC/mvc/services"
	"github.com/nwneisen/MicroservicesInGo/Section2-MVC/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userIdParam := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		jsonValue, _ := json.Marshal(&utils.Error{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
			Code:    "bad request",
		})
		resp.Write(jsonValue)
		return
	}

	user, err := services.GetUser(userId)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		jsonValue, _ := json.Marshal(&utils.Error{
			Message: err.Error(),
			Status:  http.StatusNotFound,
			Code:    "not found",
		})
		resp.Write(jsonValue)
		return
	}

	// Return user to client
	jsonValue, err := json.Marshal(user)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		jsonValue, _ := json.Marshal(&utils.Error{
			Message: err.Error(),
			Status:  http.StatusNotFound,
			Code:    "not found",
		})
		resp.Write(jsonValue)
		return
	}

	resp.Write(jsonValue)
}
