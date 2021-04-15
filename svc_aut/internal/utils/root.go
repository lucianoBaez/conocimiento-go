package utils

import (
	"os"
)

// GetEnv Get environment variable value or fallback value
func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
