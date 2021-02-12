package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func main() {
	db, _ = gorm.Open("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(
		&Product{},
		&Stock{},
		&Provider{},
		&Bill{},
		&DeliveryOrder{},
		&PurchaseOrder{},
		&TypeProd{},
		&Users{},
		&Client{}
	)


}