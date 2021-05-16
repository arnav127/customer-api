package main

import (
	"database/sql"
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

	//add user to database
	queryString := fmt.Sprintf("insert into users values ('%v', '%v', '%v', '%v', '%v') "+
		"returning id, firstname, lastname, email, phone",
		newUser.Id, newUser.FirstName, newUser.LastName, newUser.Email, newUser.Phone)
	query, err := db.Query(queryString)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		if _, err := response.Write([]byte(`{ "error" : "` + err.Error() + `" 	}`)); err != nil {
			panic(err)
		}
		return
	}
	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			panic(err)
		}
	}(query)
	var user User
	if query.Next() {
		err = query.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Phone)
		if err != nil {
			panic(err)
		}
	}
	if err := json.NewEncoder(response).Encode(user); err != nil {
		panic(err)
	}
	fmt.Println("Creating customer", user)

}
