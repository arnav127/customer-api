package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestMainAll (t *testing.T) {
	go main()
	time.Sleep(1 * time.Second)

	//createCustomer tests
	validUser := User { "12", "Arnav", "Dixit", "arnavdixit@email.com", 7413680639}
	invalidUser := User { "12", "", "Dixit", "arnavdixit@email.com", 7413680639}
	bytesRepresentation, _ := json.Marshal(validUser)
	_, err := http.Post("http://0.0.0.0:8080/v2/user", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		t.Error(err)
	}
	bytesRepresentation, _ = json.Marshal(invalidUser)
	_, err = http.Post("http://0.0.0.0:8080/v2/user", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		t.Error(err)
	}

	//searchCustomers
	resp, err := http.Get("http://0.0.0.0:8080/v2/user?email=arnavdixit@email.com&first_name=arnav")
	if err != nil {
		t.Error(err)
	}
	var testUser User
	json.NewDecoder(resp.Body).Decode(&testUser)
	//fmt.Println(testUser)
	if testUser != validUser {
		t.Error("Search failed")
	}

	//listAllCustomers
	var listusers []User
	resp, err = http.Get("http://0.0.0.0:8080/v2/users")
	if err != nil {
		t.Error(err)
	}
	json.NewDecoder(resp.Body).Decode(&listusers)
	if listusers[0] != validUser {
		t.Error("ListAll failed")
	}

	//updateCustomer
	reqUrl, _ := url.ParseRequestURI("http://0.0.0.0:8080/user")
	updatedUser := validUser
	updatedUser.FirstName = "Manish"
	updatedUser.Email = "manish@gmail.com"
	bytesRepresentation, _ = json.Marshal(updatedUser)
	req := &http.Request{
		Method:           "PATCH",
		URL:              reqUrl,
		Header:           map[string][]string {
			"Content-Type": { "application/json" },
		},
		Body:             ioutil.NopCloser(bytes.NewBuffer(bytesRepresentation)),
	}
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}

	//deleteCustomer
	reqUrl, _ = url.ParseRequestURI("http://0.0.0.0:8080/user?id=12")
	deleteReq := &http.Request{
		Method: "DELETE",
		URL: reqUrl,
		Header:           map[string][]string {
			"Content-Type": { "application/json" },
		},
	}

	_, err = http.DefaultClient.Do(deleteReq)
	if err != nil {
		t.Error(err)
	}
}
