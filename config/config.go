package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

/**
 * Config struct
 */
type Config struct {
	ApiUrl    string
	TodoCount int
}

/**
 * NewConfig creates a new Config
 * @return *Config
 */

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default values")
	}

	ApiUrl := getEnv("BASE_URL", "https://jsonplaceholder.typicode.com/todos/")
	TodoCount := getEnvAsInt("MAX_LIMIT", 20)

	return &Config{
		ApiUrl:    ApiUrl,
		TodoCount: TodoCount,
	}
}

/**
 * getEnv function
 * @param key string
 * @param defaultValue string
 * @return string
 */
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

/**
 * getEnvAsInt function
 * @param name string
 * @param defaultVal int
 * @return int
 */
func getEnvAsInt(name string, defaultVal int) int {
	if valueStr, exists := os.LookupEnv(name); exists {
		if value, err := strconv.Atoi(valueStr); err == nil {
			return value
		}
	}
	return defaultVal
}
