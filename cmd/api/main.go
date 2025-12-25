package main

import (
	"log"
	"net/http"
	"resume-analyzer/internal/config"
)

func main() {
	cfg := config.LoadConfig()
	db, err := config.ConnectDB(cfg.DatabaseUrl)
	if err != nil {
		log.Fatal("Database Error: ", err)
	}

	log.Println("Server is running on port "+ cfg.Port + ". Driver name: " + db.DriverName())
	if err := http.ListenAndServe(cfg.Port, nil); err != nil {
		log.Fatal(err)
	}
}