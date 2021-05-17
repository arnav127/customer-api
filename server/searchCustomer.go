package server

import (
	"database/sql"
	"errors"
)

func SearchCustomer(email string, firstName string) (*User, error) {

		var userResult User
		const queryString = "select id, firstname, lastname, email, phone from users where email=$1 and firstname=$2"

		//Query database and generate userResult from row returned
		row := Db.QueryRow(queryString, email, firstName)
		err := row.Scan(&userResult.Id, &userResult.FirstName, &userResult.LastName, &userResult.Email, &userResult.Phone)

		switch err {
		//row returned successfully
		case nil:
			return &userResult, nil
		//no row returned
		case sql.ErrNoRows:
			return nil, errors.New("customer not in database")
		default:
			return nil, err
		}
}
