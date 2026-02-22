package config

import (
	"SpendSmartAPI/internal/util"
	"os"

	"github.com/joho/godotenv"
)

type ConfigurationEnv struct {
	Port             int
	DatabasePort     int
	DatabaseHost     string
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
}

func GetEnvConfig() (*ConfigurationEnv, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	return &ConfigurationEnv{
		Port:             util.GetEnvInt(os.Getenv("PORT"), 8000),
		DatabasePort:     util.GetEnvInt(os.Getenv("DB_PORT"), 3306),
		DatabaseHost:     util.GetOrDefaultPrimitive(os.Getenv("DB_HOST"), "localhost"),
		DatabaseName:     util.GetOrDefaultPrimitive(os.Getenv("DB_NAME"), "spendsmart"),
		DatabaseUser:     util.GetOrDefaultPrimitive(os.Getenv("DB_USER"), "root"),
		DatabasePassword: util.GetOrDefaultPrimitive(os.Getenv("DB_PASS"), ""),
	}, nil
}
