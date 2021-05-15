package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func createCustomer(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	//Only POST requests allowed
	if request.Method != http.MethodPost {
		response.WriteHeader(http.StatusMethodNotAllowed)
		if _, err := response.Write([]byte(`{ "error": "Method not allowed" }`)); err != nil {
			panic(err)
		}
		return
	}

	// decode and check for error
	var newUser User
	if err := json.NewDecoder(request.Body).Decode(&newUser); err != nil {
		response.WriteHeader(http.StatusBadRequest)
		if _, err := response.Write([]byte(`{ "error": "Bad request" }`)); err != nil {
			panic(err)
		}
		return
	}

	//validate user details provided
	if err := validate(newUser, false); err != nil {
		response.WriteHeader(http.StatusBadRequest)
		if _, err := response.Write([]byte(`{ "error" : "` + err.Error() + `" 	}`)); err != nil {
			panic(err)
		}
		return
	}

	//add user to slice
	users = append(users, newUser)
	fmt.Println("Creating customer", newUser)

	//Return newly created user
	if err := json.NewEncoder(response).Encode(newUser); err != nil {
		panic(err)
	}

}
