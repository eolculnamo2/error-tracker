package main

import (
	"log"
//	"os"
	"github.com/eolculnamo2/error-tracker/api/src"
	"github.com/eolculnamo2/error-tracker/api/src/db"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func dbInit() {
	var err error
	// 
	// connectionString := os.Getenv("db_credentials") + os.Getenv("db_connection") + os.Getenv("db_parameters")
	connectionString := "root:root@(localhost:3306)/error-tracker?charset=utf8&parseTime=True&loc=Local"
	log.Println(connectionString)
	db.DbConnection, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic("Failed to connect to the database" + err.Error())
	}
}

func main() {
	log.Println("Server starting...")
	dbInit()
  defer db.DbConnection.Close()
	app.StartApp(db.DbConnection)
}