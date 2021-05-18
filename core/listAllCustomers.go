package core

import (
	customer_api "gitlab.com/arnavdixit/customer-api"
)

func (s *Service) ListAllCustomer() *[]customer_api.DbUser {
	//query database and store results in slice
	var user customer_api.DbUser
	var usersList []customer_api.DbUser

	const queryString = "SELECT id, firstname, lastname, email, phone FROM users order by id"
	rows, _ := s.Db.Query(queryString)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Phone)
		if err != nil {
			panic(err)
		}
		usersList = append(usersList, user)
	}
	//Return userList obtained from database
	return &usersList
}
