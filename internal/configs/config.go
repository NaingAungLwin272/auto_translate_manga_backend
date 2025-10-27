package configs

import (
	"log"
	"os"
)

type Config struct {
	Port     string
	MongoURI string
	DBName   string
}

func NewConfig() *Config {
	port := os.Getenv("PORT")
	mongoURI := os.Getenv("MONGO_URL")
	dbName := os.Getenv("DB_NAME")

	if port == "" {
		log.Fatal("Environment variable PORT is required")
	}
	if mongoURI == "" {
		log.Fatal("Environment variable MONGO_URI is required")
	}
	if dbName == "" {
		log.Fatal("Environment variable DB_NAME is required")
	}

	return &Config{
		Port:     port,
		MongoURI: mongoURI,
		DBName:   dbName,
	}
}
