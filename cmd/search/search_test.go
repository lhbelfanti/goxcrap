package search_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"goxcrap/cmd/page"
	"goxcrap/cmd/search"
)

func TestExecuteAdvanceSearch_success(t *testing.T) {
	mockLoadPage := page.MockMakeLoad(nil)
	executeAdvanceSearch := search.MakeExecuteAdvanceSearch(mockLoadPage)
	mockCriteria := search.MockCriteria()[0]

	got := executeAdvanceSearch(mockCriteria)

	assert.Nil(t, got)
}

func TestExecuteAdvanceSearch_failsWhenLoadPageThrowsError(t *testing.T) {
	err := errors.New("error while executing loadPage")
	mockLoadPage := page.MockMakeLoad(err)
	executeAdvanceSearch := search.MakeExecuteAdvanceSearch(mockLoadPage)
	mockCriteria := search.MockCriteria()[0]

	want := search.NewSearchError(search.FailedToLoadAdvanceSearchPage, err)
	got := executeAdvanceSearch(mockCriteria)

	assert.Equal(t, want, got)
}
