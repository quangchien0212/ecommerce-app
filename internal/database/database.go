package database

import (
	"fmt"

	"github.com/quangchien0212/ecommerce-app/internal/abstract"
	"github.com/quangchien0212/ecommerce-app/internal/config"
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

func NewDBClient(c *config.Config) (DBClient, error) {
	config := c.Get()
	dbHost := config.GetString("DB_HOST")
	dbUsername := config.GetString("DB_USERNAME")
	dbPassword := config.GetString("DB_PASSWORD")
	dbName := config.GetString("DB_NAME")
	dbPort := config.GetUint("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		dbHost, dbUsername, dbPassword, dbName, dbPort, "disable")

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
