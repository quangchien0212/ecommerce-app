package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/quangchien0212/ecommerce-app/internal/abstract"
	"github.com/quangchien0212/ecommerce-app/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBClient interface {
	Ready() bool
	RunMigration() error
	CloseConnection()
	abstract.Category
}

type Client struct {
	DB *gorm.DB
}

func NewDBClient() (DBClient, error) {
	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	databasePort, err := strconv.Atoi(dbPort)
	if err != nil {
		log.Fatal("Invalid DB Port")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		dbHost, dbUsername, dbPassword, dbName, databasePort, "disable")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	client := Client{DB: db}

	return client, nil
}

func (c Client) Ready() bool {
	var ready string
	result := c.DB.Raw("SELECT 1 as ready").Scan(&ready)
	if result.Error != nil {
		return false
	}
	if ready == "1" {
		return true
	}
	return false
}
func (c Client) RunMigration() error {
	err := c.DB.AutoMigrate(
		&models.Category{},
		&models.Product{},
	)
	if err != nil {
		return err
	}
	return nil
}

func (c Client) CloseConnection() {
	sqlDB, _ := c.DB.DB()
	err := sqlDB.Close()
	if err != nil {
		return
	}
}

func NewTestDBClient() (DBClient, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	client := Client{DB: db}
	return client, nil
}
