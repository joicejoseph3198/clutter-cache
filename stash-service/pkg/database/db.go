package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"joicejoseph.dev/clutter-cache/stash-service/config"
	"joicejoseph.dev/clutter-cache/stash-service/pkg/model"
)

type DB struct {
	*gorm.DB
}

func NewPostgresConnection() *DB {
	// Read database configuration from the config package
	postgresConfig := config.GetPostgresConfig()

	// Build the connection string
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		postgresConfig.Host, postgresConfig.Port, postgresConfig.User, postgresConfig.Password, postgresConfig.Name,
	)

	// Open the database connection
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	} else {
		log.Printf("database connection established")
	}

	return &DB{db}
}

// makes all neccessary migrations
func SyncDatabase(db *DB) {
	// if err := db.Migrator().DropTable(&model.Entry{}); err != nil {
	// 	log.Printf("Table already exists: %v", err)
	// }
	if err := db.AutoMigrate(&model.Entry{}); err != nil {
		log.Printf("Table already exists: %v", err)
	}
	log.Println("Database migration successful")
}
