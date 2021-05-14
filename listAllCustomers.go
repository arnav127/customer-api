package main

import (
	"encoding/json"
	"net/http"
)

func listAllCustomer(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	if request.Method != "GET" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		if _, err := response.Write([]byte(`{ "error": "Method not allowed" }`)); err != nil {
			panic(err)
		}
		return
	}

	//Return user list
	if err := json.NewEncoder(response).Encode(users); err != nil {
		panic(err)
	}
}
