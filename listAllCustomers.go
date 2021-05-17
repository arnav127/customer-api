package main

import (
	"encoding/json"
	"net/http"
	"strconv"
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

	//query limit and offset from URL, if not present take default values
	limit, offset := 10, 0
	limitQuery := request.URL.Query()["limit"]
	offsetQuery := request.URL.Query()["offset"]
	if limitQuery != nil {
		var err error
		limit, err = strconv.Atoi(limitQuery[0])
		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
			if _, err := response.Write([]byte(`{"error": "Bad Request"}`)); err != nil {
				panic(err)
			}
		}
	}
	if offsetQuery != nil {
		var err error
		offset, err = strconv.Atoi(offsetQuery[0])
		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
			if _, err := response.Write([]byte(`{"error": "Bad Request"}`)); err != nil {
				panic(err)
			}
		}
	}
	
	const queryString = "SELECT id, firstname, lastname, email, phone FROM users order by id limit $1 offset $2"
	rows, _ := db.Query(queryString, limit, offset)
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
