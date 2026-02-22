package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		"root",
		"",
		"localhost",
		3306,
		"spendsmart",
	)

	database, err := sql.Open("mysql", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	err = database.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection established")

	return database
}
