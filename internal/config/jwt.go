package config

import (
	"os"

	"github.com/joho/godotenv"
)

func GetJwtScret() string {
	godotenv.Load()
	return os.Getenv("JWT_SECRET")
}
