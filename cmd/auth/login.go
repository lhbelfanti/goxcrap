package auth

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"

	"goxcrap/cmd/scrapper"
)

const (
	pageLoaderTimeout              time.Duration = 10 * time.Second
	elementTimeout                 time.Duration = 10 * time.Second
	emailInputName                 string        = "text"
	nextButtonXPath                string        = "/html/body/div/div/div/div[1]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div/div[6]"
	passwordInputName              string        = "password"
	logInButtonXPath               string        = "/html/body/div/div/div/div[1]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div[2]/div/div[1]/div/div/div"
	usernameInputName              string        = "text"
	unusualActivityNextButtonXPath string        = "/html/body/div/div/div/div[1]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div[2]/div/div/div/div/div"
)

// Login finds de login button clicks it and then fill the email and password fields to log in the user
func Login(scrapper scrapper.Scrapper) error {
	err := scrapper.LoadPage("/i/flow/login", pageLoaderTimeout)
	if err != nil {
		return err
	}

	emailInput, err := scrapper.WaitAndRetrieveElement(selenium.ByName, emailInputName, elementTimeout)
	if err != nil {
		return NewAuthError(fmt.Sprintf(FailedToRetrieveInput, "email"), err)
	}

	err = emailInput.Click()
	if err != nil {
		return NewAuthError(fmt.Sprintf(FailedToClickInput, "email"), err)
	}

	err = emailInput.SendKeys(scrapper.Credentials.Email)
	if err != nil {
		return NewAuthError(fmt.Sprintf(FailedToFillInput, "email"), err)
	}

	nextButton, err := scrapper.Driver.FindElement(selenium.ByXPATH, nextButtonXPath)
	if err != nil {
		return NewAuthError(fmt.Sprintf(FailedToRetrieveButton, "next"), err)
	}

	err = nextButton.Click()
	if err != nil {
		return NewAuthError(fmt.Sprintf(FailedToClickButton, "next"), err)
	}

	err = ExecuteUnusualActivityFlow(scrapper)
	if err != nil {
		return err
	}

	passwordInput, err := scrapper.WaitAndRetrieveElement(selenium.ByName, passwordInputName, elementTimeout)
	if err != nil {
		return NewAuthError(fmt.Sprintf(FailedToRetrieveInput, "password"), err)
	}

	err = passwordInput.Click()
	if err != nil {
		return NewAuthError(fmt.Sprintf(FailedToClickInput, "password"), err)
	}

	err = passwordInput.SendKeys(scrapper.Credentials.Password)
	if err != nil {
		return NewAuthError(fmt.Sprintf(FailedToFillInput, "password"), err)
	}

	logInButton, err := scrapper.Driver.FindElement(selenium.ByXPATH, logInButtonXPath)
	if err != nil {
		return NewAuthError(fmt.Sprintf(FailedToRetrieveButton, "log in"), err)
	}

	err = logInButton.Click()
	if err != nil {
		return NewAuthError(fmt.Sprintf(FailedToClickButton, "log in"), err)
	}

	return nil
}

func ExecuteUnusualActivityFlow(scrapper scrapper.Scrapper) error {
	// 'There was an unusual activity in your account' flow
	usernameInput, err := scrapper.WaitAndRetrieveElement(selenium.ByName, usernameInputName, elementTimeout)
	if err != nil {
		return NewAuthError(fmt.Sprintf(FailedToRetrieveInput, "username"), err)
	}

	err = usernameInput.Click()
	if err != nil {
		return NewAuthError(fmt.Sprintf(FailedToClickInput, "username"), err)
	}

	err = usernameInput.SendKeys(scrapper.Credentials.Username)
	if err != nil {
		return NewAuthError(fmt.Sprintf(FailedToFillInput, "username"), err)
	}

	nextButton, err := scrapper.Driver.FindElement(selenium.ByXPATH, unusualActivityNextButtonXPath)
	if err != nil {
		return NewAuthError(fmt.Sprintf(FailedToRetrieveButton, "next"), err)
	}

	err = nextButton.Click()
	if err != nil {
		return NewAuthError(fmt.Sprintf(FailedToClickButton, "next"), err)
	}

	return nil
}
