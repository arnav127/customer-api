package core

import (
	"context"
	"database/sql"
	"errors"
	customer_api "gitlab.com/arnavdixit/customer-api"
)

func updateCustomerDB(Db *sql.DB, updateId string, updateUser *customer_api.DbUser) error {
	const queryString = "update users set id=$1, firstname=$2, lastname=$3, email=$4, phone=$5 where id=$6 returning *"
	const queryUpdateFirstname = "update users set firstname=$1 where id=$2"
	const queryUpdateLastname = "update users set lastname=$1 where id=$2"
	const queryUpdateEmail = "update users set email=$1 where id=$2"
	const queryUpdatePhone = "update users set phone=$1 where id=$2"
	const queryUpdateId = "update users set id=$1 where id=$2"

	//begin transaction
	ctx := context.Background()
	updateTransaction, err := Db.BeginTx(ctx, nil)
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

func (s *Service) UpdateCustomer(userId string, updateUser *customer_api.DbUser) (*customer_api.DbUser, error) {

		//
		////PUT request: update all user details
		//if request.Method == http.MethodPut {
		//
		//	//validate user details provided
		//	if err := validate(&updateUser, false); err != nil {
		//		return nil, err
		//	}
		//
		//	err := updateCustomerDB(userId, &updateUser)
		//	switch err {
		//	case nil:
		//		return &updateUser, nil
		//	default:
		//		return nil, err
		//	}
		//}

		//PATCH request: update only the values provided
		//if request.Method == http.MethodPatch {
			//validate user details provided
		if err := validate(updateUser, true); err != nil {
			return nil, err
		}

		const queryString = "select id, firstname, lastname, email, phone from users where id=$1"
		row := s.Db.QueryRow(queryString, userId)
		var currentUser customer_api.DbUser
		err := row.Scan(&currentUser.Id, &currentUser.FirstName, &currentUser.LastName, &currentUser.Email, &currentUser.Phone)
		switch err {
		//user is in db
		case nil:
			//update current user struct to reflect provided attributes
			updatedLocalUser := updateNonEmptyDetails(&currentUser, updateUser)
			//update the user in db
			err := updateCustomerDB(s.Db, userId, updatedLocalUser)
			switch err {
			case nil:
				return updatedLocalUser, nil
			default:
				return nil, err
			}

		//If customer does not exist in record return an error
		case sql.ErrNoRows:
			return nil, errors.New("customer not in database")
		default:
			return nil, err
		}
		//}
}
