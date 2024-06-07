package env_test

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"goxcrap/cmd/api/env"
)

func TestLoadVariables_success(t *testing.T) {
	_ = godotenv.Load()

	want := env.Variables{
		Email:    "email",
		Password: "password",
		Username: "username",
	}
	got := env.LoadVariables()

	assert.Equal(t, want, got)
}
