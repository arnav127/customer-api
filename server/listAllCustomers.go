package server

func ListAllCustomer() *[]User {
	//query database and store results in slice
	var user User
	var usersList []User

	const queryString = "SELECT id, firstname, lastname, email, phone FROM users order by id"
	rows, _ := Db.Query(queryString)
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
