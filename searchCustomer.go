package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func searchCustomer(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	if request.Method != "GET" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte(`{ "error": "Method not allowed" }`))
		return
	}
	searchByEmail := request.URL.Query()["email"]

	if searchByEmail != nil {
		for _, user := range users {
			if strings.ToLower(user.Email) == strings.ToLower(searchByEmail[0]) {
				json.NewEncoder(response).Encode(user)
				return
			}
		}
	} else {
		searchByFirstName := request.URL.Query()["first_name"]
		if searchByFirstName != nil {
			for _, user := range users {
				if strings.ToLower(user.FirstName) == strings.ToLower(searchByFirstName[0]) {
					json.NewEncoder(response).Encode(user)
					return
				}
			}
		} else {
			response.WriteHeader(http.StatusBadRequest)
			response.Write([]byte(`{"error": "Bad Request"}`))
			return
		}
	}
	response.WriteHeader(http.StatusNotFound)
	response.Write([]byte(`{"error": "Record not found"}`))
}
