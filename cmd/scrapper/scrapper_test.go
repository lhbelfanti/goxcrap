package scrapper_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"goxcrap/cmd/auth"
	"goxcrap/cmd/scrapper"
)

func TestInit_success(t *testing.T) {
	mockLogin := auth.MockMakeLogin(nil)

	got := scrapper.Init(mockLogin)

	assert.Nil(t, got)
}

func TestInit_failsWhenLoginThrowsError(t *testing.T) {
	want := errors.New("error while executing login")
	mockLogin := auth.MockMakeLogin(want)

	got := scrapper.Init(mockLogin)

	assert.Equal(t, got, want)
}
