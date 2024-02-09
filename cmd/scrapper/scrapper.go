package scrapper

import (
	"github.com/tebeka/selenium"
)

type (
	Scrapper struct {
		Driver      selenium.WebDriver
		Credentials LoginCredentials
		PageLoader  PageLoader
	}

	LoginCredentials struct {
		Email    string
		Password string
	}
)

// New creates a new Scrapper struct
func New(driver selenium.WebDriver, email, password string, pageLoader PageLoader) Scrapper {
	return Scrapper{
		Driver: driver,
		Credentials: LoginCredentials{
			Email:    email,
			Password: password,
		},
		PageLoader: pageLoader,
	}
}
