package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func createCustomer(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	if request.Method != "POST" {
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

	//TODO: validate user before adding to map
	//validate(newUser)

	//add user to map
	users = append(users, newUser)
	fmt.Println("Creating customer", newUser)

	//Return newly created user
	if err := json.NewEncoder(response).Encode(newUser); err != nil {
		panic(err)
	}

}
