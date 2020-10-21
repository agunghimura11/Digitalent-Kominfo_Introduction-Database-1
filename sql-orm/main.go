package main

import (
	"DB/sql-generic/config"
	"DB/sql-orm/database"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func main(){
	cfg,err := getConfig()
	if err != nil {
		log.Println(err)
		return
	}

	db,err := initDB(cfg.Database)
	if err != nil {
		log.Println(err)
		return
	}

	//insert customer data
	database.InsertCustomer(database.CustomerORM{
		FirstName: "Aril",
		Lastname: "Sudanopo",
		NpwpId: "npwp123",
		Age: 20,
		CustomerType: "Premium",
		Street: "RingRoad",
		City: "Bantul",
		State: "Indonesia",
		ZipCode: "123",
		PhoneNumber: "08882828282",
		AccountORM: []database.AccountORM{
			{
				Balance: 1000,
				AccountType: "Premium",
			},
			{
				Balance:     1000000,
				AccountType: "Deposito",
			},
		},
	}, db)

	// Get data customer
	//database.GetCustomer(db)

	// Update Database
	//database.UpdateCustomer(database.CustomerORM{
	//	FirstName: "Arielo",
	//	Age: 12,
	//	City: "Jakarta",
	//}, 1, db)

	// get first data user account
	acc, err := database.GetTranscByID(1, db)
	if err != nil {
		log.Println(err)
		return
	}
	// add transaction
	database.AddTransc(database.TranscORM{
		Date: time.Now(),
		Type: "Transfer",
		AccId: acc.IdCustomerRefer, // id customer
		Details: []database.TranscDetail{
			{
				AccFrom: "1",
				AccTo: "3",
				Amount: 10000,
			},
		},
	}, db)
}

func getConfig() (config.Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.SetConfigName("config.yml")

	if err := viper.ReadInConfig(); err != nil {
		return config.Config{}, err
	}

	var cfg config.Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}

func initDB(dbConfig config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName, dbConfig.Config)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		&database.CustomerORM{},
		&database.AccountORM{},
		&database.TranscORM{},
		&database.TranscDetail{},
	)

	log.Println("db successfully connected")

	return db, nil
}
