package database

import (
	"SpendSmartAPI/internal/config"
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	envConfig, err := config.GetEnvConfig()

	if err != nil {
		log.Println(".ENV not found, using system env")
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		envConfig.DatabaseUser,
		envConfig.DatabasePassword,
		envConfig.DatabaseHost,
		strconv.Itoa(envConfig.DatabasePort),
		envConfig.DatabaseName,
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
