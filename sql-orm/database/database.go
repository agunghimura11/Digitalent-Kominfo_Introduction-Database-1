package database

import (
	"gorm.io/gorm"
	"log"
)

type CustomerORM struct {
	ID int `gorm:"primary_key" json:"customer_id"`
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
	AccountORM []AccountORM `gorm:"ForeignKey: IdCustomerRefer" json:"account_orm"`
}

type  AccountORM struct {
	ID int `gorm:"primary_key" json:"-"`
	IdCustomerRefer int `json:"-"`
	Balance int `json:"balance"`
	AccountType string `json:"account_type"`
}

func InsertCustomer(customer CustomerORM, db *gorm.DB){
	if err := db.Create(&customer).Error; err != nil {
		log.Println("Failed to insert", err.Error())
		return
	}
	log.Println("Success Insert Data")
}

func GetCustomer (db *gorm.DB){
	var customer []CustomerORM
	// preload untuk otomatis mengambil data dari AccountORM
	if err := db.Preload("AccountORM").Find(&customer).Error;err != nil {
		log.Println("Failed",err.Error())
		return
	}
	log.Println(customer)
}
