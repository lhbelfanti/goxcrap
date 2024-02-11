package auth

import (
	"time"

	"github.com/tebeka/selenium"
	"goxcrap/cmd/element"
	"goxcrap/cmd/env"
	"goxcrap/cmd/page"
)

const (
	pageLoaderTimeout              time.Duration = 10 * time.Second
	elementTimeout                 time.Duration = 10 * time.Second
	logInPageRelativeURL           string        = "/i/flow/login"
	emailInputName                 string        = "text"
	nextButtonXPath                string        = "/html/body/div/div/div/div[1]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div/div[6]"
	passwordInputName              string        = "password"
	logInButtonXPath               string        = "/html/body/div/div/div/div[1]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div[2]/div/div[1]/div/div/div"
	usernameInputName              string        = "text"
	unusualActivityNextButtonXPath string        = "/html/body/div/div/div/div[1]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div[2]/div/div/div/div/div"
)

// Login finds de login button clicks it and then fill the email and password fields to log in the user
type Login func() error

// MakeLogin creates a new Login
func MakeLogin(envVariables env.Variables, loadPage page.Load, retrieveAndFillInput element.RetrieveAndFillInput, retrieveAndClickButton element.RetrieveAndClickButton) Login {
	return func() error {
		err := loadPage(logInPageRelativeURL, pageLoaderTimeout)
		if err != nil {
			return err
		}

		err = retrieveAndFillInput(selenium.ByName, emailInputName, "email", envVariables.Email, elementTimeout, NewAuthError)
		if err != nil {
			return err
		}

		err = retrieveAndClickButton(selenium.ByXPATH, nextButtonXPath, "next", elementTimeout, NewAuthError)
		if err != nil {
			return err
		}

		// -- 'There was an unusual activity in your account' flow
		err = retrieveAndFillInput(selenium.ByName, usernameInputName, "username", envVariables.Username, elementTimeout, NewAuthError)
		if err != nil {
			return err
		}

		err = retrieveAndClickButton(selenium.ByXPATH, unusualActivityNextButtonXPath, "next", elementTimeout, NewAuthError)
		if err != nil {
			return err
		}
		// 'There was an unusual activity in your account' flow --

		err = retrieveAndFillInput(selenium.ByName, passwordInputName, "password", envVariables.Password, elementTimeout, NewAuthError)
		if err != nil {
			return err
		}

		err = retrieveAndClickButton(selenium.ByXPATH, logInButtonXPath, "log in", elementTimeout, NewAuthError)
		if err != nil {
			return err
		}

		return nil
	}
}
