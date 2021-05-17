package server

import (
	"database/sql"
	"errors"
)

func GetCustomer(id string) (*User, error) {
	var userResult User
	const queryString = "select id, firstname, lastname, email, phone from users where id=$1"

	//Query database and generate userResult from row returned
	row := Db.QueryRow(queryString, id)
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
