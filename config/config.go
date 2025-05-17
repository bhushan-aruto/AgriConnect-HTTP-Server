package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	CallAnswerApi string
	CallFrom      string
	DatabaseUrl   string
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalln("missing or empty ", key, " env variable")
	}
	return value
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("failed to load the .env file")
	}

	return &Config{
		CallAnswerApi: getEnv("CALL_ANSWER_API"),
		CallFrom:      getEnv("CALL_FROM"),
		DatabaseUrl:   getEnv("DATABASE_URL"),
	}
}
