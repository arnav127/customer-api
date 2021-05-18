package core

import (
	"database/sql"
	"errors"
	customer_api "gitlab.com/arnavdixit/customer-api"
)

func (s *Service) GetCustomer(id string) (*customer_api.DbUser, error) {
	var userResult customer_api.DbUser
	const queryString = "select id, firstname, lastname, email, phone from users where id=$1"

	//Query database and generate userResult from row returned
	row := s.Db.QueryRow(queryString, id)
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
