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

func TestRetrieveAndClickButton_success(t *testing.T) {
	mockWebElement := new(element.MockWebElement)
	mockWebElement.On("Click").Return(nil)
	mockWaitAndRetrieve := element.MockMakeWaitAndRetrieve(mockWebElement, nil)

	retrieveAndClickButton := element.MakeRetrieveAndClickButton(mockWaitAndRetrieve)

	got := retrieveAndClickButton(selenium.ByName, "name", "element", 10*time.Minute, element.NewElementError)

	assert.Nil(t, got)
}

func TestRetrieveAndClickButton_failsWhileRetrievingTheElement(t *testing.T) {
	err := errors.New("error while retrieving the element")
	mockWaitAndRetrieve := element.MockMakeWaitAndRetrieve(nil, err)
	retrieveAndClickButton := element.MakeRetrieveAndClickButton(mockWaitAndRetrieve)

	want := element.NewElementError(fmt.Sprintf(element.FailedToRetrieveButton, "test"), err)
	got := retrieveAndClickButton(selenium.ByName, "value", "test", 10*time.Minute, element.NewElementError)

	assert.Equal(t, got, want)
}

func TestRetrieveAndClickButton_failsWhileClickingTheElement(t *testing.T) {
	err := errors.New("error while clicking the element")
	mockWebElement := new(element.MockWebElement)
	mockWebElement.On("Click").Return(err)
	mockWaitAndRetrieve := element.MockMakeWaitAndRetrieve(mockWebElement, nil)
	retrieveAndClickButton := element.MakeRetrieveAndClickButton(mockWaitAndRetrieve)

	want := element.NewElementError(fmt.Sprintf(element.FailedToClickButton, "test"), err)
	got := retrieveAndClickButton(selenium.ByName, "value", "test", 10*time.Minute, element.NewElementError)

	assert.Equal(t, got, want)
}