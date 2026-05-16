package env_test

import (
	"controle_financeiro/src/config/env"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	t.Run("should return error when DATABASE_URL is empty", func(t *testing.T) {
		t.Setenv("DATABASE_URL", "")

		err := env.Init()

		assert.Error(t, err)
		assert.Equal(t, "DATABASE_URL é obrigatório", err.Error())
	})

	t.Run("should use environment values", func(t *testing.T) {
		t.Setenv("PORT", "8080")
		t.Setenv("APP_ENV", "test")
		t.Setenv("CORS_ORIGIN", "http://localhost:5173")
		t.Setenv("CORS_METHOD", "OPTIONS,GET,PUT,DELETE,POST")
		t.Setenv("CORS_HEADER", "*")
		t.Setenv("DATABASE_URL", "test.db")

		err := env.Init()

		assert.NoError(t, err)
		assert.Equal(t, "8080", env.Port)
		assert.Equal(t, "test", env.AppEnv)
		assert.Equal(t, "http://localhost:5173", env.CorsOrigin)
		assert.Equal(t, "OPTIONS,GET,PUT,DELETE,POST", env.CorsMethod)
		assert.Equal(t, "*", env.CorsHeader)
		assert.Equal(t, "test.db", env.DatabaseURL)
	})
}
