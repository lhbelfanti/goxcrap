package search_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"goxcrap/cmd/api/search"
)

func TestCriteriaDTO_ToType_success(t *testing.T) {
	want := search.MockCriteria()
	got := search.MockCriteriaDTO().ToType()

	assert.Equal(t, want, got)
}
