package auth_test

import (
	"errors"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"goxcrap/cmd/auth"
	"goxcrap/cmd/element"
	"goxcrap/cmd/env"
	"goxcrap/cmd/page"
)

func TestMain(m *testing.M) {
	_ = godotenv.Load()
	os.Exit(m.Run())
}

func TestLogin_success(t *testing.T) {
	mockWebElement := new(element.MockWebElement)
	envVariables := env.LoadVariables()
	mockLoadPage := page.MockMakeLoad(nil)
	mockWaitAndRetrieveElement := element.MockMakeWaitAndRetrieve(mockWebElement, nil)
	mockRetrieveAndFillInput := element.MockMakeRetrieveAndFillInput(nil, "")
	mockRetrieveAndClickButton := element.MockMakeRetrieveAndClickButton(nil, "")
	login := auth.MakeLogin(envVariables, mockLoadPage, mockWaitAndRetrieveElement, mockRetrieveAndFillInput, mockRetrieveAndClickButton)

	got := login()

	assert.Nil(t, got)
}

func TestLogin_successWhenWaitAndRetrievePasswordElementThrowsError(t *testing.T) {
	err := errors.New("error while executing waitAndRetrieveElement")
	mockWebElement := new(element.MockWebElement)
	envVariables := env.LoadVariables()
	mockLoadPage := page.MockMakeLoad(nil)
	mockWaitAndRetrieveElement := element.MockMakeWaitAndRetrieve(mockWebElement, err)
	mockRetrieveAndFillInput := element.MockMakeRetrieveAndFillInput(nil, "")
	mockRetrieveAndClickButton := element.MockMakeRetrieveAndClickButton(nil, "")
	login := auth.MakeLogin(envVariables, mockLoadPage, mockWaitAndRetrieveElement, mockRetrieveAndFillInput, mockRetrieveAndClickButton)

	got := login()

	assert.Nil(t, got)
}

func TestLogin_failsWhenLoadPageThrowsError(t *testing.T) {
	want := errors.New("error while executing loadPage")
	mockWebElement := new(element.MockWebElement)
	envVariables := env.LoadVariables()
	mockLoadPage := page.MockMakeLoad(want)
	mockWaitAndRetrieveElement := element.MockMakeWaitAndRetrieve(mockWebElement, nil)
	mockRetrieveAndFillInput := element.MockMakeRetrieveAndFillInput(nil, "")
	mockRetrieveAndClickButton := element.MockMakeRetrieveAndClickButton(nil, "")
	login := auth.MakeLogin(envVariables, mockLoadPage, mockWaitAndRetrieveElement, mockRetrieveAndFillInput, mockRetrieveAndClickButton)

	got := login()

	assert.Equal(t, got, want)
}

func TestLogin_failsWhenRetrieveAndFillEmailInputThrowsError(t *testing.T) {
	want := errors.New("error while executing retrieveAndFillEmailInput")
	mockWebElement := new(element.MockWebElement)
	envVariables := env.LoadVariables()
	mockLoadPage := page.MockMakeLoad(nil)
	mockWaitAndRetrieveElement := element.MockMakeWaitAndRetrieve(mockWebElement, nil)
	mockRetrieveAndFillInput := element.MockMakeRetrieveAndFillInput(want, "email input")
	mockRetrieveAndClickButton := element.MockMakeRetrieveAndClickButton(nil, "")
	login := auth.MakeLogin(envVariables, mockLoadPage, mockWaitAndRetrieveElement, mockRetrieveAndFillInput, mockRetrieveAndClickButton)

	got := login()

	assert.Equal(t, got, want)
}

func TestLogin_failsWhenRetrieveAndClickEmailNextButtonThrowsError(t *testing.T) {
	want := errors.New("error while executing retrieveAndClickEmailNextButton")
	mockWebElement := new(element.MockWebElement)
	envVariables := env.LoadVariables()
	mockLoadPage := page.MockMakeLoad(nil)
	mockWaitAndRetrieveElement := element.MockMakeWaitAndRetrieve(mockWebElement, nil)
	mockRetrieveAndFillInput := element.MockMakeRetrieveAndFillInput(nil, "")
	mockRetrieveAndClickButton := element.MockMakeRetrieveAndClickButton(want, "email next button")
	login := auth.MakeLogin(envVariables, mockLoadPage, mockWaitAndRetrieveElement, mockRetrieveAndFillInput, mockRetrieveAndClickButton)

	got := login()

	assert.Equal(t, got, want)
}

func TestLogin_failsWhenRetrieveAndFillUsernameInputThrowsError(t *testing.T) {
	want := errors.New("error while executing retrieveAndFillUsernameInput")
	mockWebElement := new(element.MockWebElement)
	envVariables := env.LoadVariables()
	mockLoadPage := page.MockMakeLoad(nil)
	mockWaitAndRetrieveElement := element.MockMakeWaitAndRetrieve(mockWebElement, errors.New("error while executing waitAndRetrieveElement"))
	mockRetrieveAndFillInput := element.MockMakeRetrieveAndFillInput(want, "username input")
	mockRetrieveAndClickButton := element.MockMakeRetrieveAndClickButton(nil, "")
	login := auth.MakeLogin(envVariables, mockLoadPage, mockWaitAndRetrieveElement, mockRetrieveAndFillInput, mockRetrieveAndClickButton)

	got := login()

	assert.Equal(t, got, want)
}

func TestLogin_failsWhenRetrieveAndClickUsernameNextButtonThrowsError(t *testing.T) {
	want := errors.New("error while executing retrieveAndClickUsernameNextButton")
	mockWebElement := new(element.MockWebElement)
	envVariables := env.LoadVariables()
	mockLoadPage := page.MockMakeLoad(nil)
	mockWaitAndRetrieveElement := element.MockMakeWaitAndRetrieve(mockWebElement, errors.New("error while executing waitAndRetrieveElement"))
	mockRetrieveAndFillInput := element.MockMakeRetrieveAndFillInput(nil, "")
	mockRetrieveAndClickButton := element.MockMakeRetrieveAndClickButton(want, "username next button")
	login := auth.MakeLogin(envVariables, mockLoadPage, mockWaitAndRetrieveElement, mockRetrieveAndFillInput, mockRetrieveAndClickButton)

	got := login()

	assert.Equal(t, got, want)
}

func TestLogin_failsWhenRetrieveAndFillPasswordInputThrowsError(t *testing.T) {
	want := errors.New("error while executing retrieveAndFillPasswordInput")
	mockWebElement := new(element.MockWebElement)
	envVariables := env.LoadVariables()
	mockLoadPage := page.MockMakeLoad(nil)
	mockWaitAndRetrieveElement := element.MockMakeWaitAndRetrieve(mockWebElement, nil)
	mockRetrieveAndFillInput := element.MockMakeRetrieveAndFillInput(want, "password input")
	mockRetrieveAndClickButton := element.MockMakeRetrieveAndClickButton(nil, "")
	login := auth.MakeLogin(envVariables, mockLoadPage, mockWaitAndRetrieveElement, mockRetrieveAndFillInput, mockRetrieveAndClickButton)

	got := login()

	assert.Equal(t, got, want)
}

func TestLogin_failsWhenRetrieveAndClickLogInButtonThrowsError(t *testing.T) {
	want := errors.New("error while executing retrieveAndClickLogInButton")
	mockWebElement := new(element.MockWebElement)
	envVariables := env.LoadVariables()
	mockLoadPage := page.MockMakeLoad(nil)
	mockWaitAndRetrieveElement := element.MockMakeWaitAndRetrieve(mockWebElement, nil)
	mockRetrieveAndFillInput := element.MockMakeRetrieveAndFillInput(nil, "")
	mockRetrieveAndClickButton := element.MockMakeRetrieveAndClickButton(want, "log in button")
	login := auth.MakeLogin(envVariables, mockLoadPage, mockWaitAndRetrieveElement, mockRetrieveAndFillInput, mockRetrieveAndClickButton)

	got := login()

	assert.Equal(t, got, want)
}
