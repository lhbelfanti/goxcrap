package elements_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tebeka/selenium"

	"goxcrap/cmd/api/elements"
)

func TestRetrieveAndClickButton_success(t *testing.T) {
	mockButtonWebElement := new(elements.MockWebElement)
	mockButtonWebElement.On("Click").Return(nil)
	mockWaitAndRetrieve := elements.MockWaitAndRetrieve(mockButtonWebElement, nil)

	retrieveAndClickButton := elements.MakeRetrieveAndClickButton(mockWaitAndRetrieve)

	got := retrieveAndClickButton(context.Background(), selenium.ByName, "name", "element", 10*time.Minute)

	assert.Nil(t, got)
}

func TestRetrieveAndClickButton_failsWhenWaitAndRetrieveElementThrowsError(t *testing.T) {
	mockWaitAndRetrieve := elements.MockWaitAndRetrieve(nil, errors.New("error while executing waitAndRetrieveElement"))

	retrieveAndClickButton := elements.MakeRetrieveAndClickButton(mockWaitAndRetrieve)

	want := elements.FailedToRetrieveButton
	got := retrieveAndClickButton(context.Background(), selenium.ByName, "value", "test", 10*time.Minute)

	assert.Equal(t, want, got)
}

func TestRetrieveAndClickButton_failsWhenButtonClickThrowsError(t *testing.T) {
	mockButtonWebElement := new(elements.MockWebElement)
	mockButtonWebElement.On("Click").Return(errors.New("error while executing button.Click"))
	mockWaitAndRetrieve := elements.MockWaitAndRetrieve(mockButtonWebElement, nil)

	retrieveAndClickButton := elements.MakeRetrieveAndClickButton(mockWaitAndRetrieve)

	want := elements.FailedToClickButton
	got := retrieveAndClickButton(context.Background(), selenium.ByName, "value", "test", 10*time.Minute)

	assert.Equal(t, want, got)
}
