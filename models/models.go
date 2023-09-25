package models

type CustomerInfo struct {
	FirstName      string  `json:"first_name" bson:"first_name"`
	LastName       string  `json:"last_name" bson:"last_name"`
	CustomerID     string  `json:"customer_id" bson:"customer_id"`
	SSNNumber      int64   `json:"ssn_number" bson:"ssn_number"`
	PhoneNumber    int64   `json:"phone_number" bson:"phone_number"`
	Email          string  `json:"email" bson:"email"`
	DOB            DOB     `json:"dob" bson:"dob"`
	Address        Address `json:"address" bson:"address"`
	JobTitle       string  `json:"job_title" bson:"job_title"`
	AnnualIncome   float64 `json:"annual_income" bson:"annual_income"`
	InitialDeposit float64 `json:"initial_deposit" bson:"initial_deposit"`
}

type DOB struct {
	Year  int64 `json:"year" bson:"year"`
	Month int64 `json:"month" bson:"month"`
	Date  int64 `json:"date" bson:"date"`
}

type Address struct {
	StreetAddress string `json:"street_address" bson:"street_address"`
	City          string `json:"city" bson:"city"`
	State         string `json:"state" bson:"state"`
	PostalCode    string `json:"postal_code" bson:"postal_code"`
	Country       string `json:"country" bson:"country"`
}

type DBresponse struct {
	SSNNumber   int64  `json:"ssn_number" bson:"ssn_number"`
	PhoneNumber int64  `json:"phone_number" bson:"phone_number"`
	Email       string `json:"email" bson:"email"`
}
