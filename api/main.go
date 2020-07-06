package main

import (
	"fmt"
	"os"
	"github.com/eolculnamo2/error-tracker/api/src"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	fmt.Println("Server starting...")
	// jdbc:mysql://localhost:3306/icet?useTimezone=true&serverTimezone=UTC
	connectionString := os.Getenv("db_credentials") + os.Getenv("db_connection") + os.Getenv("db_parameters")
	fmt.Println(connectionString)
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic("Failed to connect to the database" + err.Error())
	}
  defer db.Close()
	app.StartApp(db)
}