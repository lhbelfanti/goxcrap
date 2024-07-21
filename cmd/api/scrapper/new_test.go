package scrapper_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"goxcrap/cmd/api/scrapper"
	"goxcrap/internal/webdriver"
)

func TestNew_success(t *testing.T) {
	mockWebDriver := new(webdriver.Mock)
	makeNew := scrapper.MakeNew()

	got := makeNew(mockWebDriver)

	assert.NotNil(t, got)
}
