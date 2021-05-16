package main

import (
	"encoding/json"
	"net/http"
)

func listAllCustomer(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	//Only GET requests allowed
	if request.Method != http.MethodGet {
		response.WriteHeader(http.StatusMethodNotAllowed)
		if _, err := response.Write([]byte(`{ "error": "Method not allowed" }`)); err != nil {
			panic(err)
		}
		return
	}

	//query database and store results in slice
	var user User
	var usersList []User
	rows, _ := db.Query("SELECT id, firstname, lastname, email, phone FROM users")
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Phone)
		if err != nil {
			panic(err)
		}
		usersList = append(usersList, user)
	}

	//Return userList obtained from database
	if err := json.NewEncoder(response).Encode(usersList); err != nil {
		panic(err)
	}
}
