package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	Port        string
	AppEnv      string
	DatabaseURL string
	CorsOrigin  string
	CorsMethod  string
	CorsHeader  string
)

func Init() error {
	_ = godotenv.Load()

	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "8080"
	}

	AppEnv = os.Getenv("APP_ENV")
	if AppEnv == "" {
		AppEnv = "development"
	}

	CorsOrigin = os.Getenv("CORS_ORIGIN")
	if CorsOrigin == "" {
		CorsOrigin = "*"
	}

	CorsMethod = os.Getenv("CORS_METHOD")
	if CorsMethod == "" {
		CorsMethod = "*"
	}

	CorsHeader = os.Getenv("CORS_HEADER")
	if CorsHeader == "" {
		CorsHeader = "*"
	}

	DatabaseURL = os.Getenv("DATABASE_URL")
	if DatabaseURL == "" {
		return fmt.Errorf("DATABASE_URL é obrigatório")
	}

	return nil
}
