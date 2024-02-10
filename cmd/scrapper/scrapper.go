package scrapper

import (
	"github.com/tebeka/selenium"
)

type (
	Scrapper struct {
		Driver                 selenium.WebDriver
		Credentials            LoginCredentials
		TakeScreenshot         TakeScreenshot
		LoadPage               LoadPage
		WaitAndRetrieveElement WaitAndRetrieveElement
	}

	LoginCredentials struct {
		Email    string
		Password string
		Username string
	}
)

// New creates a new Scrapper struct
func New(
	driver selenium.WebDriver,
	email, password, username string,
	takeScreenshot TakeScreenshot,
	pageLoader LoadPage,
	waitAndRetrieveElement WaitAndRetrieveElement,
) Scrapper {

	return Scrapper{
		Driver: driver,
		Credentials: LoginCredentials{
			Email:    email,
			Password: password,
			Username: username,
		},
		TakeScreenshot:         takeScreenshot,
		LoadPage:               pageLoader,
		WaitAndRetrieveElement: waitAndRetrieveElement,
	}
}
