package auth_test

import (
	"errors"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"goxcrap/cmd/api/auth"
	"goxcrap/cmd/api/elements"
	"goxcrap/cmd/api/env"
	"goxcrap/cmd/api/page"
)

func TestMain(m *testing.M) {
	_ = godotenv.Load()
	os.Exit(m.Run())
}

func TestLogin_success(t *testing.T) {
	mockWebElement := new(elements.MockWebElement)
	mockEnvVariables := env.LoadVariables()
	mockLoadPage := page.MockLoad(nil)
	mockWaitAndRetrieveElement := elements.MockWaitAndRetrieve(mockWebElement, nil)
	mockRetrieveAndFillInput := elements.MockRetrieveAndFillInput(nil, "")
	mockRetrieveAndClickButton := elements.MockRetrieveAndClickButton(nil, "")

	login := auth.MakeLogin(mockEnvVariables, mockLoadPage, mockWaitAndRetrieveElement, mockRetrieveAndFillInput, mockRetrieveAndClickButton)

	got := login()

	assert.Nil(t, got)
}

func TestLogin_successWhenWaitAndRetrievePasswordElementThrowsError(t *testing.T) {
	err := errors.New("error while executing waitAndRetrieveElement")
	mockWebElement := new(elements.MockWebElement)
	mockEnvVariables := env.LoadVariables()
	mockLoadPage := page.MockLoad(nil)
	mockWaitAndRetrieveElement := elements.MockWaitAndRetrieve(mockWebElement, err)
	mockRetrieveAndFillInput := elements.MockRetrieveAndFillInput(nil, "")
	mockRetrieveAndClickButton := elements.MockRetrieveAndClickButton(nil, "")

	login := auth.MakeLogin(mockEnvVariables, mockLoadPage, mockWaitAndRetrieveElement, mockRetrieveAndFillInput, mockRetrieveAndClickButton)

	got := login()

	assert.Nil(t, got)
}

func TestLogin_failsWhenLoadPageThrowsError(t *testing.T) {
	want := errors.New("error while executing loadPage")
	mockWebElement := new(elements.MockWebElement)
	mockEnvVariables := env.LoadVariables()
	mockLoadPage := page.MockLoad(want)
	mockWaitAndRetrieveElement := elements.MockWaitAndRetrieve(mockWebElement, nil)
	mockRetrieveAndFillInput := elements.MockRetrieveAndFillInput(nil, "")
	mockRetrieveAndClickButton := elements.MockRetrieveAndClickButton(nil, "")

	login := auth.MakeLogin(mockEnvVariables, mockLoadPage, mockWaitAndRetrieveElement, mockRetrieveAndFillInput, mockRetrieveAndClickButton)

	got := login()

	assert.Equal(t, got, want)
}

func TestLogin_failsWhenRetrieveAndFillEmailInputThrowsError(t *testing.T) {
	want := errors.New("error while executing retrieveAndFillEmailInput")
	mockWebElement := new(elements.MockWebElement)
	mockEnvVariables := env.LoadVariables()
	mockLoadPage := page.MockLoad(nil)
	mockWaitAndRetrieveElement := elements.MockWaitAndRetrieve(mockWebElement, nil)
	mockRetrieveAndFillInput := elements.MockRetrieveAndFillInput(want, "email input")
	mockRetrieveAndClickButton := elements.MockRetrieveAndClickButton(nil, "")

	login := auth.MakeLogin(mockEnvVariables, mockLoadPage, mockWaitAndRetrieveElement, mockRetrieveAndFillInput, mockRetrieveAndClickButton)

	got := login()

	assert.Equal(t, got, want)
}

func TestLogin_failsWhenRetrieveAndClickEmailNextButtonThrowsError(t *testing.T) {
	want := errors.New("error while executing retrieveAndClickEmailNextButton")
	mockWebElement := new(elements.MockWebElement)
	mockEnvVariables := env.LoadVariables()
	mockLoadPage := page.MockLoad(nil)
	mockWaitAndRetrieveElement := elements.MockWaitAndRetrieve(mockWebElement, nil)
	mockRetrieveAndFillInput := elements.MockRetrieveAndFillInput(nil, "")
	mockRetrieveAndClickButton := elements.MockRetrieveAndClickButton(want, "email next button")

	login := auth.MakeLogin(mockEnvVariables, mockLoadPage, mockWaitAndRetrieveElement, mockRetrieveAndFillInput, mockRetrieveAndClickButton)

	got := login()

	assert.Equal(t, got, want)
}

func TestLogin_failsWhenRetrieveAndFillUsernameInputThrowsError(t *testing.T) {
	want := errors.New("error while executing retrieveAndFillUsernameInput")
	mockWebElement := new(elements.MockWebElement)
	mockEnvVariables := env.LoadVariables()
	mockLoadPage := page.MockLoad(nil)
	mockWaitAndRetrieveElement := elements.MockWaitAndRetrieve(mockWebElement, errors.New("error while executing waitAndRetrieveElement"))
	mockRetrieveAndFillInput := elements.MockRetrieveAndFillInput(want, "username input")
	mockRetrieveAndClickButton := elements.MockRetrieveAndClickButton(nil, "")

	login := auth.MakeLogin(mockEnvVariables, mockLoadPage, mockWaitAndRetrieveElement, mockRetrieveAndFillInput, mockRetrieveAndClickButton)

	got := login()

	assert.Equal(t, got, want)
}

func TestLogin_failsWhenRetrieveAndClickUsernameNextButtonThrowsError(t *testing.T) {
	want := errors.New("error while executing retrieveAndClickUsernameNextButton")
	mockWebElement := new(elements.MockWebElement)
	mockEnvVariables := env.LoadVariables()
	mockLoadPage := page.MockLoad(nil)
	mockWaitAndRetrieveElement := elements.MockWaitAndRetrieve(mockWebElement, errors.New("error while executing waitAndRetrieveElement"))
	mockRetrieveAndFillInput := elements.MockRetrieveAndFillInput(nil, "")
	mockRetrieveAndClickButton := elements.MockRetrieveAndClickButton(want, "username next button")

	login := auth.MakeLogin(mockEnvVariables, mockLoadPage, mockWaitAndRetrieveElement, mockRetrieveAndFillInput, mockRetrieveAndClickButton)

	got := login()

	assert.Equal(t, got, want)
}

func TestLogin_failsWhenRetrieveAndFillPasswordInputThrowsError(t *testing.T) {
	want := errors.New("error while executing retrieveAndFillPasswordInput")
	mockWebElement := new(elements.MockWebElement)
	mockEnvVariables := env.LoadVariables()
	mockLoadPage := page.MockLoad(nil)
	mockWaitAndRetrieveElement := elements.MockWaitAndRetrieve(mockWebElement, nil)
	mockRetrieveAndFillInput := elements.MockRetrieveAndFillInput(want, "password input")
	mockRetrieveAndClickButton := elements.MockRetrieveAndClickButton(nil, "")

	login := auth.MakeLogin(mockEnvVariables, mockLoadPage, mockWaitAndRetrieveElement, mockRetrieveAndFillInput, mockRetrieveAndClickButton)

	got := login()

	assert.Equal(t, got, want)
}

func TestLogin_failsWhenRetrieveAndClickLogInButtonThrowsError(t *testing.T) {
	want := errors.New("error while executing retrieveAndClickLogInButton")
	mockWebElement := new(elements.MockWebElement)
	mockEnvVariables := env.LoadVariables()
	mockLoadPage := page.MockLoad(nil)
	mockWaitAndRetrieveElement := elements.MockWaitAndRetrieve(mockWebElement, nil)
	mockRetrieveAndFillInput := elements.MockRetrieveAndFillInput(nil, "")
	mockRetrieveAndClickButton := elements.MockRetrieveAndClickButton(want, "log in button")

	login := auth.MakeLogin(mockEnvVariables, mockLoadPage, mockWaitAndRetrieveElement, mockRetrieveAndFillInput, mockRetrieveAndClickButton)

	got := login()

	assert.Equal(t, got, want)
}
