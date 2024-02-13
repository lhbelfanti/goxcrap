package env_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"goxcrap/cmd/env"
)

func TestMain(m *testing.M) {
	_ = godotenv.Load()
	os.Exit(m.Run())
}

func TestLoadVariables_success(t *testing.T) {
	want := env.Variables{
		Email:    "email",
		Password: "password",
		Username: "username",
	}
	got := env.LoadVariables()

	assert.Equal(t, got, want)
}
