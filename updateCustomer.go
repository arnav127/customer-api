package main

import (
	"encoding/json"
	"net/http"
)

func updateCustomer(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	if request.Method != "PUT" && request.Method != "PATCH" {
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

	if request.Method == "PATCH" {
		for idx, _ := range users {
			if users[idx].Id == updateUser.Id {
				if updateUser.FirstName != "" {
					users[idx].FirstName = updateUser.FirstName
				}
				if updateUser.LastName != "" {
					users[idx].LastName = updateUser.LastName
				}
				if updateUser.Email != "" {
					users[idx].Email = updateUser.Email
				}
				if updateUser.Phone != 0 {
					users[idx].Phone = updateUser.Phone
				}
				response.WriteHeader(http.StatusOK)
				responseEncoder.Encode(users[idx])
				return
			}
		}
	}

	users = append(users, updateUser)
	responseEncoder.Encode(updateUser)
}
