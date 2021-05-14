package main

import (
	"net/http"
)

func deleteCustomer(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	if request.Method != "GET" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte(`{ "error": "Method not allowed" }`))
		return
	}
	deleteID := request.URL.Query()["id"]
	if deleteID != nil {
		for idx, user := range users {
			if user.Id == deleteID[0] {
				totalUsers := len(users)
				users[totalUsers-1], users[idx] = users[idx], users[totalUsers-1]
				users = users[:totalUsers-1]
				response.WriteHeader(http.StatusNoContent)
				return
			}
		}
	} else {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "error": "Bad Request" }`))
	}
	response.WriteHeader(http.StatusNotFound)
	response.Write([]byte(`{ "error": "Customer does not exist" }`))

}
