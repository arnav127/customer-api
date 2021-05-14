package main

type User struct {
	id string `json:"id"`
	firstName string `json:"first_name"`
	lastName string `json:"last_name"`
	email string `json:"email"`
	phone int `json:"phone"`
}

var users map[string]User