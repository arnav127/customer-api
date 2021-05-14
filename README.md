# Customer API

Basic CRUD Customer API in Golang


customer properties
|- ID
|- FIRST NAME
|- LAST NAME
|- EMAIL
|- PHONE

Domain: localhost:8080

Customers application
|- Create customer
|- Update customer
|- Delete customer
|- List All customers
|- Search customer

Create Customer
Request
POST http://localhost:8080/create HTTP/1.1
Body: {"id":"0987654", "first_name":"shikhar"}
Response
OK / Error

=> Search Customer with email: example@gmail.com and first_name: example
Request
GET http://localhost:8080/search?email=example@gmail.com&first_name=example

GET http://localhost:8080/unknown -> 404 Not Found
