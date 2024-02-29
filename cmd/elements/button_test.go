package elements_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"

	"goxcrap/cmd/elements"
)

func TestRetrieveAndClickButton_success(t *testing.T) {
	mockWebElement := new(elements.MockWebElement)
	mockWebElement.On("Click").Return(nil)
	mockWaitAndRetrieve := elements.MockWaitAndRetrieve(mockWebElement, nil)

	retrieveAndClickButton := elements.MakeRetrieveAndClickButton(mockWaitAndRetrieve)

	got := retrieveAndClickButton(selenium.ByName, "name", "element", 10*time.Minute)

	assert.Nil(t, got)
}

func TestRetrieveAndClickButton_failsWhenWaitAndRetrieveElementThrowsError(t *testing.T) {
	err := errors.New("error while executing waitAndRetrieveElement")
	mockWaitAndRetrieve := elements.MockWaitAndRetrieve(nil, err)
	retrieveAndClickButton := elements.MakeRetrieveAndClickButton(mockWaitAndRetrieve)

	want := elements.FailedToRetrieveButton
	got := retrieveAndClickButton(selenium.ByName, "value", "test", 10*time.Minute)

	assert.Equal(t, want, got)
}

func TestRetrieveAndClickButton_failsWhenButtonClickThrowsError(t *testing.T) {
	err := errors.New("error while executing button.Click")
	mockWebElement := new(elements.MockWebElement)
	mockWebElement.On("Click").Return(err)
	mockWaitAndRetrieve := elements.MockWaitAndRetrieve(mockWebElement, nil)
	retrieveAndClickButton := elements.MakeRetrieveAndClickButton(mockWaitAndRetrieve)

	want := elements.FailedToClickButton
	got := retrieveAndClickButton(selenium.ByName, "value", "test", 10*time.Minute)

	assert.Equal(t, want, got)
}
