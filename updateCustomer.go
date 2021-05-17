package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func updateCustomerDB(updateId string, updateUser User) error {
	const queryString = "update users set id=$1, firstname=$2, lastname=$3, email=$4, phone=$5 where id=$6 returning *"
	const queryUpdateFirstname = "update users set firstname=$1 where id=$2"
	const queryUpdateLastname = "update users set lastname=$1 where id=$2"
	const queryUpdateEmail = "update users set email=$1 where id=$2"
	const queryUpdatePhone = "update users set phone=$1 where id=$2"
	const queryUpdateId = "update users set id=$1 where id=$2"

	//begin transaction
	ctx := context.Background()
	updateTransaction, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	//update firstname
	_, err = updateTransaction.ExecContext(ctx, queryUpdateFirstname, updateUser.FirstName, updateId)
	if err != nil {
		updateTransaction.Rollback()
		return err
	}
	//update lastname
	_, err = updateTransaction.ExecContext(ctx, queryUpdateLastname, updateUser.LastName, updateId)
	if err != nil {
		updateTransaction.Rollback()
		return err
	}
	//update email
	_, err = updateTransaction.ExecContext(ctx, queryUpdateEmail, updateUser.Email, updateId)
	if err != nil {
		updateTransaction.Rollback()
		return err
	}
	//update phone
	_, err = updateTransaction.ExecContext(ctx, queryUpdatePhone, updateUser.Phone, updateId)
	if err != nil {
		updateTransaction.Rollback()
		return err
	}
	//update id
	_, err = updateTransaction.ExecContext(ctx, queryUpdateId, updateUser.Id, updateId)
	if err != nil {
		updateTransaction.Rollback()
		return err
	}
	//commit transaction
	err = updateTransaction.Commit()
	if err != nil {
		return err
	}
	return nil
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

			err := updateCustomerDB(updateID[0], updateUser)
			switch err {
			case nil:
				if err := json.NewEncoder(response).Encode(updateUser); err != nil {
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
				err := updateCustomerDB(updateID[0], updatedLocalUser)
				switch err {
				case nil:
					response.WriteHeader(http.StatusOK)
					if err := json.NewEncoder(response).Encode(updatedLocalUser); err != nil {
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
