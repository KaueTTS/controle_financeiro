package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	Port               string
	AppEnv             string
	DatabaseURL        string
	FrontendCorsOrigin string
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

	FrontendCorsOrigin = os.Getenv("FRONTEND_CORS_ORIGIN")
	if FrontendCorsOrigin == "" {
		FrontendCorsOrigin = "*"
	}

	DatabaseURL = os.Getenv("DATABASE_URL")
	if DatabaseURL == "" {
		return fmt.Errorf("DATABASE_URL é obrigatório")
	}

	return nil
}
