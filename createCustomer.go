package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func createCustomer(response http.ResponseWriter, request *http.Request) {

	if request.Method != "POST" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte(`{ "error": "Method not allowed" }`))
		return
	}

	// decode and check for error
	var newUser User
	if err := json.NewDecoder(request.Body).Decode(&newUser); err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "error": "Bad request" }`))
		return
	}

	//TODO: validate user before adding to map
	//validate(newUser)

	//add user to map
	users = append(users, newUser)
	fmt.Println("Creating customer", newUser)

	//Return newly created user
	response.Header().Set("content-type", "application/json")
	json.NewEncoder(response).Encode(newUser)

}
