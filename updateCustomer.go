package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func updateCustomerDB(updateId string, updateUser User) (User, error) {
	var updatedUser User
	const queryString = "update users set id=$1, firstname=$2, lastname=$3, email=$4, phone=$5 where id=$6 returning *"

	//query db update user and scan returned row into updatedUser
	row := db.QueryRow(queryString, updateUser.Id, updateUser.FirstName, updateUser.LastName, updateUser.Email, updateUser.Phone, updateId)
	err := row.Scan(&updatedUser.Id, &updatedUser.FirstName, &updatedUser.LastName, &updatedUser.Email, &updatedUser.Phone)
	switch err {
	//user successfully updated
	case nil:
		return updatedUser, nil
	//query not successful, return error
	default:
		return User{}, err
	}
}

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

	updateID := request.URL.Query()["id"]
	if updateID != nil {

		//PUT request: update all user details
		if request.Method == http.MethodPut {

			//validate user details provided
			if err := validate(updateUser, false); err != nil {
				fmt.Println("Sending back error")
				response.WriteHeader(http.StatusBadRequest)
				if _, err := response.Write([]byte(`{ "error" : "` + err.Error() + `" 	}`)); err != nil {
					panic(err)
				}
				return
			}

			updatedUser, err := updateCustomerDB(updateID[0], updateUser)
			switch err {
			case nil:
				if err := json.NewEncoder(response).Encode(updatedUser); err != nil {
					panic(err)
				}
			default:
				response.WriteHeader(http.StatusBadRequest)
				if _, err := response.Write([]byte(`{ "error" : "` + err.Error() + `" 	}`)); err != nil {
					panic(err)
				}

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

			const queryString = "select id, firstname, lastname, email, phone from users where id=$1"
			row := db.QueryRow(queryString, updateID[0])
			var currentUser User
			err := row.Scan(&currentUser.Id, &currentUser.FirstName, &currentUser.LastName, &currentUser.Email, &currentUser.Phone)
			switch err {
			//user is in db
			case nil:
				//update current user struct to reflect provided attributes
				updatedLocalUser := updateNonEmptyDetails(currentUser, updateUser)
				//update the user in db
				updatedDBUser, err := updateCustomerDB(updateID[0], updatedLocalUser)
				switch err {
				case nil:
					response.WriteHeader(http.StatusOK)
					if err := json.NewEncoder(response).Encode(updatedDBUser); err != nil {
						panic(err)
					}
				default:
					response.WriteHeader(http.StatusBadRequest)
					if _, err := response.Write([]byte(`{ "error" : "` + err.Error() + `" 	}`)); err != nil {
						panic(err)
					}
				}

			//If customer does not exist in record return an error
			case sql.ErrNoRows:
				response.WriteHeader(http.StatusNotFound)
				if _, err := response.Write([]byte(`{ "error": "Customer does not exist" }`)); err != nil {
					panic(err)
				}
				return

			default:
				response.WriteHeader(http.StatusNotFound)
				if _, err := response.Write([]byte(`{ "error": "` + err.Error() + `" }`)); err != nil {
					panic(err)
				}
				return
			}
		}
	} else {
		response.WriteHeader(http.StatusBadRequest)
		if _, err := response.Write([]byte(`{ "error": "Bad Request"`)); err != nil {
			panic(err)
		}
	}
}
