package auth

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/tebeka/selenium"

	"goxcrap/cmd/api/elements"
	"goxcrap/cmd/api/page"
	"goxcrap/internal/log"
)

const (
	logInPageRelativeURL string = "/i/flow/login"

	emailInputName                 string = "text"
	passwordInputName              string = "password"
	usernameInputName              string = "text"
	nextButtonXPath                string = "/html/body/div/div/div/div[1]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div/div/div/button[2]"
	logInButtonXPath               string = "/html/body/div/div/div/div[1]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div[2]/div/div[1]/div/div/button"
	unusualActivityNextButtonXPath string = "/html/body/div/div/div/div[1]/div/div/div/div/div/div/div[2]/div[2]/div/div/div[2]/div[2]/div[2]/div/div/div/button/div"
)

// Login finds de login button clicks it and then fill the email and password fields to log in the user
type Login func(ctx context.Context) error

// MakeLogin creates a new Login
func MakeLogin(loadPage page.Load, waitAndRetrieveElement elements.WaitAndRetrieve, retrieveAndFillInput elements.RetrieveAndFillInput, retrieveAndClickButton elements.RetrieveAndClickButton) Login {
	pageLoaderTimeoutValue, _ := strconv.Atoi(os.Getenv("LOGIN_PAGE_TIMEOUT"))
	pageLoaderTimeout := time.Duration(pageLoaderTimeoutValue) * time.Second
	elementTimeoutValue, _ := strconv.Atoi(os.Getenv("LOGIN_ELEMENTS_TIMEOUT"))
	elementTimeout := time.Duration(elementTimeoutValue) * time.Second
	passwordElementTimeoutValue, _ := strconv.Atoi(os.Getenv("LOGIN_PASSWORD_TIMEOUT"))
	passwordElementTimeout := time.Duration(passwordElementTimeoutValue) * time.Second

	return func(ctx context.Context) error {

		err := loadPage(ctx, logInPageRelativeURL, pageLoaderTimeout)
		if err != nil {
			return err
		}

		err = retrieveAndFillInput(ctx, selenium.ByName, emailInputName, "email input", os.Getenv("EMAIL"), elementTimeout)
		if err != nil {
			return err
		}

		err = retrieveAndClickButton(ctx, selenium.ByXPATH, nextButtonXPath, "email next button", elementTimeout)
		if err != nil {
			return err
		}

		_, err = waitAndRetrieveElement(ctx, selenium.ByName, passwordInputName, passwordElementTimeout)
		if err != nil {
			// If the password input element is not rendered it is probably because the flow
			// 'There was an unusual activity in your account', was triggered. So we need to fill the username input,
			// and then we can fill the password input
			err = retrieveAndFillInput(ctx, selenium.ByName, usernameInputName, "username input", os.Getenv("USERNAME"), elementTimeout)
			if err != nil {
				return err
			}

			err = retrieveAndClickButton(ctx, selenium.ByXPATH, unusualActivityNextButtonXPath, "username next button", elementTimeout)
			if err != nil {
				return err
			}
		}

		err = retrieveAndFillInput(ctx, selenium.ByName, passwordInputName, "password input", os.Getenv("PASSWORD"), elementTimeout)
		if err != nil {
			return err
		}

		err = retrieveAndClickButton(ctx, selenium.ByXPATH, logInButtonXPath, "log in button", elementTimeout)
		if err != nil {
			return err
		}

		log.Info(ctx, "Log In completed")
		return nil
	}
}
