package main

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