package element_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"

	"goxcrap/cmd/element"
)

func TestRetrieveAndFillInput_success(t *testing.T) {
	mockWebElement := new(element.MockWebElement)
	mockWebElement.On("Click").Return(nil)
	mockWebElement.On("SendKeys", "input").Return(nil)
	mockWaitAndRetrieve := element.MockMakeWaitAndRetrieve(mockWebElement, nil)
	retrieveAndClickButton := element.MakeRetrieveAndFillInput(mockWaitAndRetrieve)

	got := retrieveAndClickButton(selenium.ByName, "name", "element", "input", 10*time.Minute, element.NewElementError)

	assert.Nil(t, got)
}

func TestRetrieveAndFillInput_failsWhileRetrievingTheElement(t *testing.T) {
	err := errors.New("error while retrieving the element")
	mockWaitAndRetrieve := element.MockMakeWaitAndRetrieve(nil, err)
	retrieveAndClickButton := element.MakeRetrieveAndFillInput(mockWaitAndRetrieve)

	want := element.NewElementError(fmt.Sprintf(element.FailedToRetrieveInput, "test"), err)
	got := retrieveAndClickButton(selenium.ByName, "name", "test", "input", 10*time.Minute, element.NewElementError)

	assert.Equal(t, got, want)
}

func TestRetrieveAndFillInput_failsWhileClickingTheElement(t *testing.T) {
	err := errors.New("error while clicking the element")
	mockWebElement := new(element.MockWebElement)
	mockWebElement.On("Click").Return(err)
	mockWaitAndRetrieve := element.MockMakeWaitAndRetrieve(mockWebElement, nil)
	retrieveAndClickButton := element.MakeRetrieveAndFillInput(mockWaitAndRetrieve)

	want := element.NewElementError(fmt.Sprintf(element.FailedToClickInput, "test"), err)
	got := retrieveAndClickButton(selenium.ByName, "name", "test", "input", 10*time.Minute, element.NewElementError)

	assert.Equal(t, got, want)
}

func TestRetrieveAndFillInput_failsWhileFillingTheElement(t *testing.T) {
	err := errors.New("error while filling the element")
	mockWebElement := new(element.MockWebElement)
	mockWebElement.On("Click").Return(nil)
	mockWebElement.On("SendKeys", "input").Return(err)
	mockWaitAndRetrieve := element.MockMakeWaitAndRetrieve(mockWebElement, nil)
	retrieveAndClickButton := element.MakeRetrieveAndFillInput(mockWaitAndRetrieve)

	want := element.NewElementError(fmt.Sprintf(element.FailedToFillInput, "test"), err)
	got := retrieveAndClickButton(selenium.ByName, "name", "test", "input", 10*time.Minute, element.NewElementError)

	assert.Equal(t, got, want)
}
