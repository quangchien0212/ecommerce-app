package tests

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/quangchien0212/ecommerce-app/internal/database"
)

func Setup() (database.DBClient, error) {
	slog.Info("Initializing Setup..")
	testDb, err := database.NewTestDBClient()
	if err != nil {
		panic(err)
	}
	err = testDb.RunMigration()
	if err != nil {
		return nil, err
	}
	return testDb, err
}

func Teardown(testDB database.DBClient) {
	slog.Info("Removing TestDB...")
	testDB.CloseConnection()
	currentDirectory, _ := os.Getwd()
	filePath := filepath.Join(currentDirectory, "test.db")
	err := os.Remove(filePath)
	if err != nil {
		return
	}
}
