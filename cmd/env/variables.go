package env

import (
	"os"

	"github.com/joho/godotenv"
)

// Variables contains all the environment variables
type Variables struct {
	Email    string
	Password string
	Username string
}

// LoadVariables loads environment variables
func LoadVariables() (Variables, error) {
	err := godotenv.Load()
	if err != nil {
		return Variables{}, err
	}

	return Variables{
		Email:    os.Getenv("EMAIL"),
		Password: os.Getenv("PASSWORD"),
		Username: os.Getenv("USERNAME"),
	}, nil
}
