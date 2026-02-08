package criteria_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lhbelfanti/goxcrap/cmd/api/search/criteria"
)

func TestCriteriaDTO_ToType_success(t *testing.T) {
	want := criteria.MockCriteria()
	got := criteria.MockCriteriaDTO().ToType()

	assert.Equal(t, want, got)
}
