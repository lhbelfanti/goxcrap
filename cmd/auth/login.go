package auth

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"

	"goxcrap/cmd/scrapper"
)

const (
	pageLoaderTimeout time.Duration = 10 * time.Second
)

// Login finds de login button clicks it and then fill the email and password fields to log in the user
func Login(scrapper scrapper.Scrapper) error {
	driver := scrapper.Driver
	err := scrapper.PageLoader("https://twitter.com", pageLoaderTimeout)
	if err != nil {
		return err
	}

	loginButton, err := driver.FindElements(selenium.ByXPATH, "/html/body/div/div/div/div[2]/main/div/div/div[1]/div[1]/div/div[3]/div[5]/a")
	if err != nil {
		return err
	}

	fmt.Print(loginButton)

	return nil
}
