package search_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"goxcrap/cmd/page"
	"goxcrap/cmd/search"
)

func TestExecuteAdvanceSearch_success(t *testing.T) {
	mockLoadPage := page.MockLoad(nil)
	mockCriteria := search.MockCriteria()[0]

	executeAdvanceSearch := search.MakeExecuteAdvanceSearch(mockLoadPage)

	got := executeAdvanceSearch(mockCriteria)

	assert.Nil(t, got)
}

func TestExecuteAdvanceSearch_failsWhenLoadPageThrowsError(t *testing.T) {
	err := errors.New("error while executing loadPage")
	mockLoadPage := page.MockLoad(err)
	mockCriteria := search.MockCriteria()[0]

	executeAdvanceSearch := search.MakeExecuteAdvanceSearch(mockLoadPage)

	want := search.NewSearchError(search.FailedToLoadAdvanceSearchPage, err)
	got := executeAdvanceSearch(mockCriteria)

	assert.Equal(t, want, got)
}
