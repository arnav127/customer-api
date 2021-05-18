package core

import (
	customer_api "gitlab.com/arnavdixit/customer-api"
)

func (s *Service) CreateCustomer(newUser *customer_api.DbUser) (*customer_api.DbUser, error) {

	//validate user details provided
	if err := validate(newUser, false); err != nil {
		return &customer_api.DbUser{}, err
	}

	var user customer_api.DbUser
	const queryString = "insert into users values ($1, $2, $3, $4, $5) returning id, firstname, lastname, email, phone"

	//add user to database and scan row returned
	rowReturned := s.Db.QueryRow(queryString, newUser.Id, newUser.FirstName, newUser.LastName, newUser.Email, newUser.Phone)
	err := rowReturned.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Phone)

	switch err {
	//row returned successfully
	case nil:
		return &user, nil
	//error creating row
	default:
		return nil, err
	}

}
