package database

import (
	"database/sql"
	"log"
)

type Customer struct {
	CustomerId int `json:"customer_id"`
	FirstName string `json:"first_name"`
	Lastname string `json:"last_name"`
	NpwpId string `json:"npwp_id"`
	Age int `json:"age"`
	CustomerType string `json:"customer_type"`
	Street string `json:"street"`
	City string `json:"city"`
	State string `json:"state"`
	ZipCode string `json:"zip_code"`
	PhoneNumber string `json:"phone_number"`
}
// Create Database
func InsertCustomer (customer Customer, db *sql.DB){
	_, err := db.Exec("insert into customers (first_name, last_name, npwp_id, age, customer_type, street, city, state, zip_code, phone_number) values(?,?,?,?,?,?,?,?,?,?)",
		customer.FirstName,
		customer.Lastname,
		customer.NpwpId,
		customer.Age,
		customer.CustomerType,
		customer.Street,
		customer.City,
		customer.State,
		customer.ZipCode,
		customer.PhoneNumber,
	)

	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("insert data success !")
}

func GetCustomers (db *sql.DB) {
	rows, err := db.Query("select * from custumers")
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []Customer

	for rows.Next(){
		var each = Customer{}
		var err = rows.Scan(
			&each.CustomerId,
			&each.FirstName,
			&each.Lastname,
			&each.NpwpId,
			&each.Age,
			&each.CustomerType,
			&each.Street,
			&each.City,
			&each.State,
			&each.ZipCode,
			&each.PhoneNumber,
			)
		if err !=  nil {
			log.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	log.Println(result)
}