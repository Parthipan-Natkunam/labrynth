package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load .env file...")
	}
	return os.Getenv(key)
}

func getAPIPort() int {
	port, err := strconv.Atoi(getEnvVariable("PORT"))
	if err != nil {
		log.Fatalf("ENV PORT INT Parse failed")
	}
	return port
}

var (
	//PORT API Server Port
	PORT = getAPIPort()
)
