package scrapper_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"goxcrap/cmd/api/scrapper"
	"goxcrap/internal/http"
	"goxcrap/internal/webdriver"
)

func TestNew_success(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockHTTPClient := new(http.MockHTTPClient)
	makeNew := scrapper.MakeNew(mockHTTPClient, true)

	got := makeNew(mockWebDriver)

	assert.NotNil(t, got)
}
