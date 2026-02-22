package main

import (
	"SpendSmartAPI/internal/config"
	"SpendSmartAPI/internal/handler"
	"SpendSmartAPI/internal/infra/database"
	"SpendSmartAPI/internal/infra/database/mysql"
	"SpendSmartAPI/internal/routes"
	"SpendSmartAPI/internal/usecase"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	envConfig, err := config.GetEnvConfig()

	if err != nil {
		log.Println(".ENV not found, using system env")
	}

	connection := database.GetConnection()
	port := strconv.Itoa(envConfig.Port)

	userRepository := mysql.NewUserMySQLRepository(connection)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)

	r := gin.Default()

	routes.RegisterUserRoutes(r, userHandler)

	r.Run(":" + port)
}
