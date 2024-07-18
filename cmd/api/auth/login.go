package auth

import (
	"log/slog"
	"time"

	"github.com/tebeka/selenium"

	"goxcrap/cmd/api/elements"
	"goxcrap/cmd/api/env"
	"goxcrap/cmd/api/page"
)

const (
	pageLoaderTimeout      time.Duration = 10 * time.Second
	elementTimeout         time.Duration = 10 * time.Second
	passwordElementTimeout time.Duration = 5 * time.Second

	logInPageRelativeURL string = "/i/flow/login"

	emailInputName                 string = "text"
	passwordInputName              string = "password"
	usernameInputName              string = "text"
	nextButtonXPath                string = "/html/body/div/div/div/div[1]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div/button[2]/div"
	logInButtonXPath               string = "/html/body/div/div/div/div[1]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div[2]/div/div[1]/div/div/button/div"
	unusualActivityNextButtonXPath string = "/html/body/div/div/div/div[1]/div[2]/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div[2]/div/div/div/button/div"
)

// Login finds de login button clicks it and then fill the email and password fields to log in the user
type Login func() error

// MakeLogin creates a new Login
func MakeLogin(envVariables env.Variables, loadPage page.Load, waitAndRetrieveElement elements.WaitAndRetrieve, retrieveAndFillInput elements.RetrieveAndFillInput, retrieveAndClickButton elements.RetrieveAndClickButton) Login {
	return func() error {
		err := loadPage(logInPageRelativeURL, pageLoaderTimeout)
		if err != nil {
			slog.Error(err.Error())
			return err
		}

		err = retrieveAndFillInput(selenium.ByName, emailInputName, "email input", envVariables.Email, elementTimeout)
		if err != nil {
			slog.Error(err.Error())
			return err
		}

		err = retrieveAndClickButton(selenium.ByXPATH, nextButtonXPath, "email next button", elementTimeout)
		if err != nil {
			slog.Error(err.Error())
			return err
		}

		_, err = waitAndRetrieveElement(selenium.ByName, passwordInputName, passwordElementTimeout)
		if err != nil {
			// If the password input element is not rendered it is probably because the flow
			// 'There was an unusual activity in your account', was triggered. So we need to fill the username input,
			// and then we can fill the password input
			err = retrieveAndFillInput(selenium.ByName, usernameInputName, "username input", envVariables.Username, elementTimeout)
			if err != nil {
				slog.Error(err.Error())
				return err
			}

			err = retrieveAndClickButton(selenium.ByXPATH, unusualActivityNextButtonXPath, "username next button", elementTimeout)
			if err != nil {
				slog.Error(err.Error())
				return err
			}
		}

		err = retrieveAndFillInput(selenium.ByName, passwordInputName, "password input", envVariables.Password, elementTimeout)
		if err != nil {
			slog.Error(err.Error())
			return err
		}

		err = retrieveAndClickButton(selenium.ByXPATH, logInButtonXPath, "log in button", elementTimeout)
		if err != nil {
			slog.Error(err.Error())
			return err
		}

		slog.Info("Log In completed")
		return nil
	}
}
