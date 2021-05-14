package main

import (
	"encoding/json"
	"net/http"
)

func listAllCustomer(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	if request.Method != "GET" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte(`{ "error": "Method not allowed" }`))
		return
	}

	//Return user list
	json.NewEncoder(response).Encode(users)
}
