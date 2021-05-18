package core

import (
	"database/sql"
	"errors"
	customer_api "gitlab.com/arnavdixit/customer-api"
)

func (s *Service) SearchCustomer(email string, firstName string) (*customer_api.DbUser, error) {

		var userResult customer_api.DbUser
		const queryString = "select id, firstname, lastname, email, phone from users where email=$1 and firstname=$2"

		//Query database and generate userResult from row returned
		row := s.Db.QueryRow(queryString, email, firstName)
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
