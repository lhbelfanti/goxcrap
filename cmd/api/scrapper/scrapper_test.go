package scrapper_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lhbelfanti/goxcrap/cmd/api/scrapper"
	"github.com/lhbelfanti/goxcrap/internal/http"
	"github.com/lhbelfanti/goxcrap/internal/webdriver"
)

func TestNew_success(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	mockHTTPClient := new(http.MockHTTPClient)
	makeNew := scrapper.MakeNew(mockHTTPClient, true)

	got := makeNew(mockWebDriver)

	assert.NotNil(t, got)
}
