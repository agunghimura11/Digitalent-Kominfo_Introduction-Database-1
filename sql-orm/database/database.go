package database

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
	"time"
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

type TranscORM struct {
	ID int `gorm:"primary_key" json:"-"`
	Date time.Time `json:"date"`
	Type string `json:"type"` // transfer, debit, kredit
	AccId int `json:"acc_id"`
	Details []TranscDetail `gorm:"ForeignKey: IdTranscDetail" json:"transc_detail"`
}

type TranscDetail struct {
	ID int `gorm:"primary_key" json:"-"`
	IdTranscDetail int `json:"-"`
	AccFrom string `json:"acc_from"`
	AccTo string `json:"acc_to"`
	Amount int `json:"amount"`
}

func GetTranscByID(id int,  db *gorm.DB) (*AccountORM,error) {
	var acc AccountORM
	if err := db.Model(&AccountORM{}).Where(&AccountORM{ID:id}).Find(&acc).Error; err != nil{
		log.Println("Failed",err.Error())
		return nil,errors.Errorf("invalid Token")
	}
	return &acc, nil
}

func AddTransc(tranc TranscORM, db *gorm.DB){
	if err := db.Create(&tranc).Error; err != nil {
		log.Println("Transaction Failed", err.Error())
		return
	}
	log.Println("Transaction success!")
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

func DeleteCustomer (id int, db *gorm.DB){
	var customer CustomerORM
	if err := db.Where(&CustomerORM{ID: id}).Delete(&customer).Error; err != nil {
		log.Println("Failde to delete", err.Error())
		return
	}

	log.Println("success delete data")
}

func UpdateCustomer(customer CustomerORM, id int, db *gorm.DB){
	if err := db.Model(&CustomerORM{}).Where(&CustomerORM{ID:id}).Updates(customer).Error; err != nil {
		log.Println("Failed to update", err.Error())
		return
	}
	log.Println("SUccess")
}