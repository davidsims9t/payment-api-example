package graph

import (
	"fmt"
	"log"
	"os"
	"subgraph-posts/graph/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *gorm.DB
}

func NewResolver() *Resolver {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using system env")
	}

	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	pw := os.Getenv("PW")
	port := os.Getenv("PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pw, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Run auto-migrations (optional, based on your schema)
	db.AutoMigrate(&models.Post{})

	return &Resolver{DB: db}
}
