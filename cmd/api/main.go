package main

import (
	"log"
	"net/http"
	"resume-analyzer/internal/config"
	"resume-analyzer/internal/handlers"
	"resume-analyzer/internal/repository/postgres"
	v1 "resume-analyzer/internal/routes/v1"
	"resume-analyzer/internal/services"
)

func main() {
	cfg := config.LoadConfig()
	db, err := config.ConnectDB(cfg.DatabaseUrl)

	userRepo := postgres.NewUserRepository(db)

	authService := services.NewAuthService(userRepo)

	authHandler := handlers.NewAuthServiceHandler(authService)

	router := v1.NewV1Router(authHandler)

	if err != nil {
		log.Fatal("Database Error: ", err)
	}

	log.Println("Server is running on port "+ cfg.Port + ". Driver name: " + db.DriverName())
	if err := http.ListenAndServe(cfg.Port, router); err != nil {
		log.Fatal(err)
	}
}