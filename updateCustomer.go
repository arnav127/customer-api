package main

import (
	"encoding/json"
	"net/http"
)

func updateCustomer(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	//Only PUT and PATCH requests allowed
	if request.Method != "PUT" && request.Method != "PATCH" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		if _, err := response.Write([]byte(`{ "error": "Method not allowed" }`)); err != nil {
			panic(err)
		}
		return
	}

	//Decode and check for error
	var updateUser User
	if err := json.NewDecoder(request.Body).Decode(&updateUser); err != nil {
		response.WriteHeader(http.StatusBadRequest)
		if _, err := response.Write([]byte(`{ "error": "Bad request" }`)); err != nil {
			panic(err)
		}
		return
	}

	//validate user details provided
	if err := validate(updateUser); err != nil {
		if _, err := response.Write([]byte(`{ "error" : "` + err.Error() + `" 	}`)); err != nil {
			panic(err)
		}
		return
	}

	responseEncoder := json.NewEncoder(response)
	//PUT request: update all user details
	if request.Method == "PUT" {
		for idx, _ := range users {
			if users[idx].Id == updateUser.Id {
				users[idx] = updateUser
				if err := responseEncoder.Encode(users[idx]); err != nil {
					panic(err)
				}
				return
			}
		}

		//create user if does not exist
		users = append(users, updateUser)
		if err := responseEncoder.Encode(updateUser); err != nil {
			panic(err)
		}
	}

	//PATCH request: update only the values provided
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
				if err := responseEncoder.Encode(users[idx]); err != nil {
					panic(err)
				}
				return
			}
		}
	}


}
