package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file, proceeding with environment variables set in the system.")
	}
	log.Printf("MONGODB_URI: %s", os.Getenv("MONGODB_URI"))
	log.Printf("SECRET_KEY: %s", os.Getenv("SECRET_KEY"))
}

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file, proceeding with environment variables set in the system.")
	}
	return os.Getenv("MONGODB_URI")
}

func EnvSecretKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file, proceeding with environment variables set in the system.")
	}
	return os.Getenv("SECRET_KEY")
}
