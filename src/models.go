package main

import "time"

type Customer struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CustomerVisit struct {
	Id         int       `json:"id"`
	CustomerId int       `json:"customerId"`
	VisitDate  time.Time `json:"visitDate"`
}
