package scrapper_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"goxcrap/cmd/auth"
	"goxcrap/cmd/scrapper"
	"goxcrap/cmd/search"
)

func TestExecute_success(t *testing.T) {
	mockLogin := auth.MockMakeLogin(nil)
	mockCriteria := search.MockCriteria()
	mockGetSearchCriteria := search.MockMakeGetSearchCriteria(mockCriteria)

	got := scrapper.Execute(mockLogin, mockGetSearchCriteria)

	assert.Nil(t, got)
}

func TestExecute_failsWhenLoginThrowsError(t *testing.T) {
	want := errors.New("error while executing login")
	mockLogin := auth.MockMakeLogin(want)
	mockCriteria := search.MockCriteria()
	mockGetSearchCriteria := search.MockMakeGetSearchCriteria(mockCriteria)

	got := scrapper.Execute(mockLogin, mockGetSearchCriteria)

	assert.Equal(t, want, got)
}
