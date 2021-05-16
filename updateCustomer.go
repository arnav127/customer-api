package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func updateCustomerDB(updateId string, updateUser User) (User, error){
	queryString := fmt.Sprintf("update users set id='%v', firstname='%v', lastname='%v', " +
		"email='%v', phone=%v where id='%v' returning *",
		updateUser.Id, updateUser.FirstName, updateUser.LastName, updateUser.Email, updateUser.Phone, updateId)
	query, err := db.Query(queryString)
	if err != nil {
		//panic(err)
		return User{}, err
	}
	var updatedUser User
	if query.Next() {
		err = query.Scan(&updatedUser.Id, &updatedUser.FirstName, &updatedUser.LastName, &updatedUser.Email, &updatedUser.Phone)
		if err != nil {
			panic(err)
		}
	}
	return updatedUser, nil
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

			updatedUser, err := updateCustomerDB(updateID[0], updateUser)
			if err != nil {
				response.WriteHeader(http.StatusBadRequest)
				if _, err := response.Write([]byte(`{ "error" : "` + err.Error() + `" 	}`)); err != nil {
					panic(err)
				}
				return
			}
			if err := responseEncoder.Encode(updatedUser); err != nil {
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

			queryString := fmt.Sprintf("select id, firstname, lastname, email, phone from users where id='%s'", updateID[0])
			query, err := db.Query(queryString)
			if err != nil {
				panic(err)
			}
			defer query.Close()
			var currentUser User
			if query.Next() {
				err = query.Scan(&currentUser.Id, &currentUser.FirstName, &currentUser.LastName, &currentUser.Email, &currentUser.Phone)
				if err != nil {
					panic(err)
				}
			} else {
				response.WriteHeader(http.StatusNotFound)
				if _, err := response.Write([]byte(`{ "error": "Customer does not exist" }`)); err != nil {
					panic(err)
				}
				return
			}
			if updateUser.Id != "" {
				currentUser.Id = updateUser.Id
			}
			if updateUser.FirstName != "" {
				currentUser.FirstName = updateUser.FirstName
			}
			if updateUser.LastName != "" {
				currentUser.LastName = updateUser.LastName
			}
			if updateUser.Email != "" {
				currentUser.Email = updateUser.Email
			}
			if updateUser.Phone != 0 {
				currentUser.Phone = updateUser.Phone
			}

			response.WriteHeader(http.StatusOK)
			updatedUser, err := updateCustomerDB(updateID[0], currentUser)
			if err != nil {
				response.WriteHeader(http.StatusBadRequest)
				if _, err := response.Write([]byte(`{ "error" : "` + err.Error() + `" 	}`)); err != nil {
					panic(err)
				}
				return
			}
			if err := responseEncoder.Encode(updatedUser); err != nil {
				panic(err)
			}
			return
		}
		//If customer does not exist in record return an error
		} else {
		response.WriteHeader(http.StatusBadRequest)
		if _, err := response.Write([]byte(`{ "error": "Bad Request"`)); err != nil {
			panic(err)
		}
	}
}
