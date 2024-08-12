package search_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"goxcrap/cmd/api/page"
	"goxcrap/cmd/api/search"
	"goxcrap/cmd/api/search/criteria"
)

func TestExecuteAdvanceSearch_success(t *testing.T) {
	mockLoadPage := page.MockLoad(nil)
	mockCriteria := criteria.MockCriteria()

	executeAdvanceSearch := search.MakeExecuteAdvanceSearch(mockLoadPage)

	got := executeAdvanceSearch(context.Background(), mockCriteria)

	assert.Nil(t, got)
}

func TestExecuteAdvanceSearch_failsWhenLoadPageThrowsError(t *testing.T) {
	err := errors.New("error while executing loadPage")
	mockLoadPage := page.MockLoad(err)
	mockCriteria := criteria.MockCriteria()

	executeAdvanceSearch := search.MakeExecuteAdvanceSearch(mockLoadPage)

	want := search.FailedToLoadAdvanceSearchPage
	got := executeAdvanceSearch(context.Background(), mockCriteria)

	assert.Equal(t, want, got)
}
