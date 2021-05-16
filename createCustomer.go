package main

import (
	"encoding/json"
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

	var user User
	const queryString = "insert into users values ($1, $2, $3, $4, $5) returning id, firstname, lastname, email, phone"

	//add user to database and scan row returned
	rowReturned := db.QueryRow(queryString, newUser.Id, newUser.FirstName, newUser.LastName, newUser.Email, newUser.Phone)
	err := rowReturned.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Phone)

	switch err {
	//row returned successfully
	case nil:
		if err := json.NewEncoder(response).Encode(user); err != nil {
			panic(err)
		}
	//error creating row
	default:
		response.WriteHeader(http.StatusBadRequest)
		if _, err := response.Write([]byte(`{ "error" : "Could not create customer" }`)); err != nil {
			panic(err)
		}
	}

}
