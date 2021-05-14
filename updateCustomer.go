package main

import (
	"encoding/json"
	"net/http"
)

func updateCustomer(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	if request.Method != "PUT"  {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte(`{ "error": "Method not allowed" }`))
		return
	}

	var updateUser User
	if err := json.NewDecoder(request.Body).Decode(&updateUser); err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "error": "Bad request" }`))
		return
	}
	responseEncoder := json.NewEncoder(response)
	if request.Method == "PUT" {
		for idx, _ := range users {
			if users[idx].Id == updateUser.Id {
				users[idx] = updateUser
				responseEncoder.Encode(users[idx])
				return
			}
		}
	}

	users = append(users, updateUser)
	responseEncoder.Encode(updateUser)
}
