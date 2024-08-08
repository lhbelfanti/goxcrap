package log_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"goxcrap/internal/log"
)

func TestParam_success(t *testing.T) {
	want := struct {
		Key   string
		Value string
	}{"key", "value"}

	got := log.Param(want.Key, want.Value)

	assert.Equal(t, want.Key, got.Key)
	assert.Equal(t, want.Value, got.Value)
}
