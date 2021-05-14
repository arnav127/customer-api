package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func searchCustomer(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	//Only GET requests allowed
	if request.Method != "GET" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		if _, err := response.Write([]byte(`{ "error": "Method not allowed" }`)); err != nil {
			panic(err)
		}
		return
	}

	//Query email and first_name from URL
	searchByEmail := request.URL.Query()["email"]
	searchByFirstName := request.URL.Query()["first_name"]
	if searchByEmail != nil && searchByFirstName != nil {
		for _, user := range users {
			//Ignoring string case while comparing
			if strings.ToLower(user.Email) == strings.ToLower(searchByEmail[0]) &&
				strings.ToLower(user.FirstName) == strings.ToLower(searchByFirstName[0]) {
				//Return user if found
				if err := json.NewEncoder(response).Encode(user); err != nil {
					panic(err)
				}
				return
			}
		}
	} else {
		//Both email and first_name cannot be queried from URL
		response.WriteHeader(http.StatusBadRequest)
		if _, err := response.Write([]byte(`{"error": "Bad Request"}`)); err != nil {
			panic(err)
		}
		return
	}

	//Return NotFound if user not in slice
	response.WriteHeader(http.StatusNotFound)
	if _, err := response.Write([]byte(`{"error": "Record not found"}`)); err != nil {
		panic(err)
	}
}
