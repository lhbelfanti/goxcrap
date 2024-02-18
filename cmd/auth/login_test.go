package auth_test

import (
	"errors"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"goxcrap/cmd/auth"
	"goxcrap/cmd/elements"
	"goxcrap/cmd/env"
	"goxcrap/cmd/page"
)

func TestMain(m *testing.M) {
	_ = godotenv.Load()
	os.Exit(m.Run())
}

func TestLogin_success(t *testing.T) {
	mockWebElement := new(elements.MockWebElement)
	mockEnvVariables := env.LoadVariables()
	mockLoadPage := page.MockMakeLoad(nil)
	mockWaitAndRetrieveElement := elements.MockMakeWaitAndRetrieve(mockWebElement, nil)
	mockRetrieveAndFillInput := elements.MockMakeRetrieveAndFillInput(nil, "")
	mockRetrieveAndClickButton := elements.MockMakeRetrieveAndClickButton(nil, "")

	login := auth.MakeLogin(mockEnvVariables, mockLoadPage, mockWaitAndRetrieveElement, mockRetrieveAndFillInput, mockRetrieveAndClickButton)

	got := login()

	assert.Nil(t, got)
}

func TestLogin_successWhenWaitAndRetrievePasswordElementThrowsError(t *testing.T) {
	err := errors.New("error while executing waitAndRetrieveElement")
	mockWebElement := new(elements.MockWebElement)
	mockEnvVariables := env.LoadVariables()
	mockLoadPage := page.MockMakeLoad(nil)
	mockWaitAndRetrieveElement := elements.MockMakeWaitAndRetrieve(mockWebElement, err)
	mockRetrieveAndFillInput := elements.MockMakeRetrieveAndFillInput(nil, "")
	mockRetrieveAndClickButton := elements.MockMakeRetrieveAndClickButton(nil, "")

	login := auth.MakeLogin(mockEnvVariables, mockLoadPage, mockWaitAndRetrieveElement, mockRetrieveAndFillInput, mockRetrieveAndClickButton)

	got := login()

	assert.Nil(t, got)
}

func TestLogin_failsWhenLoadPageThrowsError(t *testing.T) {
	want := errors.New("error while executing loadPage")
	mockWebElement := new(elements.MockWebElement)
	mockEnvVariables := env.LoadVariables()
	mockLoadPage := page.MockMakeLoad(want)
	mockWaitAndRetrieveElement := elements.MockMakeWaitAndRetrieve(mockWebElement, nil)
	mockRetrieveAndFillInput := elements.MockMakeRetrieveAndFillInput(nil, "")
	mockRetrieveAndClickButton := elements.MockMakeRetrieveAndClickButton(nil, "")

	login := auth.MakeLogin(mockEnvVariables, mockLoadPage, mockWaitAndRetrieveElement, mockRetrieveAndFillInput, mockRetrieveAndClickButton)

	got := login()

	assert.Equal(t, got, want)
}

func TestLogin_failsWhenRetrieveAndFillEmailInputThrowsError(t *testing.T) {
	want := errors.New("error while executing retrieveAndFillEmailInput")
	mockWebElement := new(elements.MockWebElement)
	mockEnvVariables := env.LoadVariables()
	mockLoadPage := page.MockMakeLoad(nil)
	mockWaitAndRetrieveElement := elements.MockMakeWaitAndRetrieve(mockWebElement, nil)
	mockRetrieveAndFillInput := elements.MockMakeRetrieveAndFillInput(want, "email input")
	mockRetrieveAndClickButton := elements.MockMakeRetrieveAndClickButton(nil, "")

	login := auth.MakeLogin(mockEnvVariables, mockLoadPage, mockWaitAndRetrieveElement, mockRetrieveAndFillInput, mockRetrieveAndClickButton)

	got := login()

	assert.Equal(t, got, want)
}

func TestLogin_failsWhenRetrieveAndClickEmailNextButtonThrowsError(t *testing.T) {
	want := errors.New("error while executing retrieveAndClickEmailNextButton")
	mockWebElement := new(elements.MockWebElement)
	mockEnvVariables := env.LoadVariables()
	mockLoadPage := page.MockMakeLoad(nil)
	mockWaitAndRetrieveElement := elements.MockMakeWaitAndRetrieve(mockWebElement, nil)
	mockRetrieveAndFillInput := elements.MockMakeRetrieveAndFillInput(nil, "")
	mockRetrieveAndClickButton := elements.MockMakeRetrieveAndClickButton(want, "email next button")

	login := auth.MakeLogin(mockEnvVariables, mockLoadPage, mockWaitAndRetrieveElement, mockRetrieveAndFillInput, mockRetrieveAndClickButton)

	got := login()

	assert.Equal(t, got, want)
}

func TestLogin_failsWhenRetrieveAndFillUsernameInputThrowsError(t *testing.T) {
	want := errors.New("error while executing retrieveAndFillUsernameInput")
	mockWebElement := new(elements.MockWebElement)
	mockEnvVariables := env.LoadVariables()
	mockLoadPage := page.MockMakeLoad(nil)
	mockWaitAndRetrieveElement := elements.MockMakeWaitAndRetrieve(mockWebElement, errors.New("error while executing waitAndRetrieveElement"))
	mockRetrieveAndFillInput := elements.MockMakeRetrieveAndFillInput(want, "username input")
	mockRetrieveAndClickButton := elements.MockMakeRetrieveAndClickButton(nil, "")

	login := auth.MakeLogin(mockEnvVariables, mockLoadPage, mockWaitAndRetrieveElement, mockRetrieveAndFillInput, mockRetrieveAndClickButton)

	got := login()

	assert.Equal(t, got, want)
}

func TestLogin_failsWhenRetrieveAndClickUsernameNextButtonThrowsError(t *testing.T) {
	want := errors.New("error while executing retrieveAndClickUsernameNextButton")
	mockWebElement := new(elements.MockWebElement)
	mockEnvVariables := env.LoadVariables()
	mockLoadPage := page.MockMakeLoad(nil)
	mockWaitAndRetrieveElement := elements.MockMakeWaitAndRetrieve(mockWebElement, errors.New("error while executing waitAndRetrieveElement"))
	mockRetrieveAndFillInput := elements.MockMakeRetrieveAndFillInput(nil, "")
	mockRetrieveAndClickButton := elements.MockMakeRetrieveAndClickButton(want, "username next button")

	login := auth.MakeLogin(mockEnvVariables, mockLoadPage, mockWaitAndRetrieveElement, mockRetrieveAndFillInput, mockRetrieveAndClickButton)

	got := login()

	assert.Equal(t, got, want)
}

func TestLogin_failsWhenRetrieveAndFillPasswordInputThrowsError(t *testing.T) {
	want := errors.New("error while executing retrieveAndFillPasswordInput")
	mockWebElement := new(elements.MockWebElement)
	mockEnvVariables := env.LoadVariables()
	mockLoadPage := page.MockMakeLoad(nil)
	mockWaitAndRetrieveElement := elements.MockMakeWaitAndRetrieve(mockWebElement, nil)
	mockRetrieveAndFillInput := elements.MockMakeRetrieveAndFillInput(want, "password input")
	mockRetrieveAndClickButton := elements.MockMakeRetrieveAndClickButton(nil, "")

	login := auth.MakeLogin(mockEnvVariables, mockLoadPage, mockWaitAndRetrieveElement, mockRetrieveAndFillInput, mockRetrieveAndClickButton)

	got := login()

	assert.Equal(t, got, want)
}

func TestLogin_failsWhenRetrieveAndClickLogInButtonThrowsError(t *testing.T) {
	want := errors.New("error while executing retrieveAndClickLogInButton")
	mockWebElement := new(elements.MockWebElement)
	mockEnvVariables := env.LoadVariables()
	mockLoadPage := page.MockMakeLoad(nil)
	mockWaitAndRetrieveElement := elements.MockMakeWaitAndRetrieve(mockWebElement, nil)
	mockRetrieveAndFillInput := elements.MockMakeRetrieveAndFillInput(nil, "")
	mockRetrieveAndClickButton := elements.MockMakeRetrieveAndClickButton(want, "log in button")

	login := auth.MakeLogin(mockEnvVariables, mockLoadPage, mockWaitAndRetrieveElement, mockRetrieveAndFillInput, mockRetrieveAndClickButton)

	got := login()

	assert.Equal(t, got, want)
}
