package main

import (
	"log"
	"log/slog"

	"github.com/quangchien0212/ecommerce-app/internal/config"
	"github.com/quangchien0212/ecommerce-app/internal/database"
	"github.com/quangchien0212/ecommerce-app/internal/server"
)

type User struct {
  Name  string `json:"name" xml:"name"`
  Email string `json:"email" xml:"email"`
}

func main() {
	slog.Info("Initializing....")
	c := config.NewConfig()
	db, err := database.NewDBClient(c)
	if db.Ready() {
		slog.Info("The database is ready")
	} else {
		slog.Info("The database is not ready")
	}
	if err != nil {
		log.Fatalf("Failed DB Startup: %s\n", err)
		return
	}

	err = db.RunMigration()
	if err != nil {
		log.Fatalf("Migration Failed: %s\n", err)
		return
	}
	service := server.NewServer(db)
	log.Fatal(service.Start())
}