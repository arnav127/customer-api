package main

import (
	"errors"
	"regexp"
)

// User Struct to store user details
type User struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     int    `json:"phone"`
}

// users slice to store all users
var users []User

// validate : check if the user struct received is valid against set params
func validate(user User, allowEmpty bool) error {
	const minNameLength, maxNameLength = 3, 20
	const emailRegexString = "^[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+$"
	var emailRegex = regexp.MustCompile(emailRegexString)

	if !(allowEmpty && user.Email == "") {
		if len(user.Email) < 5 || !emailRegex.MatchString(user.Email) {
			return errors.New("invalid email")
		}
	}

	if !(allowEmpty && user.FirstName == "") {
		if len(user.FirstName) < minNameLength || len(user.FirstName) > maxNameLength {
			return errors.New("first_name should be between 3 and 20 characters")
		}
	}

	if !(allowEmpty && user.LastName == "") {
			if len(user.LastName) < minNameLength || len(user.LastName) > maxNameLength {
			return errors.New("last_name should be between 3 and 20 characters")
		}
	}

	if !(allowEmpty && user.Phone == 0) {
		if user.Phone < 1000000000 || user.Phone > 9999999999 {
			return errors.New("invalid phone no")
		}
	}

	if !(allowEmpty && user.Id == "") {
		if user.Id == "" {
			return errors.New("id cannot be empty")
		}
	}
	return nil
}
