package env_test

import (
	"controle_financeiro/src/config/env"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	t.Run("should return error when DATABASE_URL is empty", func(t *testing.T) {
		t.Setenv("PORT", "")
		t.Setenv("APP_ENV", "")
		t.Setenv("FRONTEND_CORS_ORIGIN", "")
		t.Setenv("DATABASE_URL", "")

		err := env.Init()

		assert.Error(t, err)
		assert.Equal(t, "DATABASE_URL é obrigatório", err.Error())
	})

	t.Run("should use default values", func(t *testing.T) {
		t.Setenv("PORT", "")
		t.Setenv("APP_ENV", "")
		t.Setenv("FRONTEND_CORS_ORIGIN", "")
		t.Setenv("DATABASE_URL", "test.db")

		err := env.Init()

		assert.NoError(t, err)
		assert.Equal(t, "8080", env.Port)
		assert.Equal(t, "development", env.AppEnv)
		assert.Equal(t, "*", env.FrontendCorsOrigin)
		assert.Equal(t, "test.db", env.DatabaseURL)
	})

	t.Run("should use environment values", func(t *testing.T) {
		t.Setenv("PORT", "3000")
		t.Setenv("APP_ENV", "test")
		t.Setenv("FRONTEND_CORS_ORIGIN", "http://localhost:5173")
		t.Setenv("DATABASE_URL", "custom.db")

		err := env.Init()

		assert.NoError(t, err)
		assert.Equal(t, "3000", env.Port)
		assert.Equal(t, "test", env.AppEnv)
		assert.Equal(t, "http://localhost:5173", env.FrontendCorsOrigin)
		assert.Equal(t, "custom.db", env.DatabaseURL)
	})
}
