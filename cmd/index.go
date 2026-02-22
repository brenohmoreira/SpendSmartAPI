package main

import (
	"SpendSmartAPI/internal/handler"
	"SpendSmartAPI/internal/infra/database"
	"SpendSmartAPI/internal/infra/database/mysql"
	"SpendSmartAPI/internal/usecase"
	"log"
	"net/http"
)

func main() {
	connection := database.GetConnection()

	// Implementação com MySQL do repositório UserRepository
	userRepository := mysql.NewUserMySQLRepository(connection)

	// NewUserUseCase espera UserRepository,
	// mas temos *UserMySQLRepository.
	// Porém *UserMySQLRepository implementa UserRepository
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)

	// Routes
	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", userHandler.Create)
	mux.HandleFunc("GET /users", userHandler.FindAll)
	mux.HandleFunc("GET /users/id", userHandler.FindById)

	log.Println("🚀 Server running on :8080")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal(err)
	}
}
