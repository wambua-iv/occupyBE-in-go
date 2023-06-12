package database

import (
	"log"

	"github.com/wambua-iv/occupyBE-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// database connection string
// format "host=localhost user=postgresUser password=postgresPassword dbname=databaseName port=defaultPort:5432 sslmode=disable"
const DSN = "add database connection string format"

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDB() {
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database \n", err.Error())
	}
	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations")

	db.AutoMigrate(&models.User{}, &models.Property{})

	Database = DbInstance{Db: db}
}
