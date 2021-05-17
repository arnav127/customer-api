package server

import (
	"errors"
)

func DeleteCustomer(id *string) error {
	//var user User
	const queryString = "delete from users where id=$1 returning *"
	query, err := Db.Exec(queryString, id)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := query.RowsAffected()
	if err != nil {
		return err
	}
	switch rowsAffected {
	case 1:
		return nil
	default:
		return errors.New("user not present in database")
	}
}
