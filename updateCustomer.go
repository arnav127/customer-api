package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func updateCustomer(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	//Only PUT and PATCH requests allowed
	if request.Method != http.MethodPut && request.Method != http.MethodPatch {
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



	responseEncoder := json.NewEncoder(response)
	//PUT request: update all user details
	if request.Method == http.MethodPut {
		//validate user details provided
		if err := validate(updateUser, false); err != nil {
			fmt.Println("Sending back error")
			response.WriteHeader(http.StatusBadRequest)
			if _, errr := response.Write([]byte(`{ "error" : "` + err.Error() + `" 	}`)); errr != nil {
				panic(errr)
			}
			return
		}
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
	if request.Method == http.MethodPatch {
		//validate user details provided
		if err := validate(updateUser, true); err != nil {
			if _, err := response.Write([]byte(`{ "error" : "` + err.Error() + `" 	}`)); err != nil {
				panic(err)
			}
			return
		}
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

		//If customer does not exist in record return an error
		response.WriteHeader(http.StatusNotFound)
		if _, err := response.Write([]byte(`{ "error": "Customer does not exist" }`)); err != nil {
			panic(err)
		}
	}


}
