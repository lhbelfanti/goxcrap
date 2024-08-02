package env

import "os"

// LoadVariables loads environment variables
func LoadVariables() Variables {
	return Variables{
		Email:       os.Getenv("EMAIL"),
		Password:    os.Getenv("PASSWORD"),
		Username:    os.Getenv("USERNAME"),
		AHBCCDomain: os.Getenv("AHBCC_DOMAIN"),
	}
}
