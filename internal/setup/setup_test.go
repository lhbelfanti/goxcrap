package setup_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"goxcrap/internal/setup"
)

func TestInit_success(t *testing.T) {
	want := "test"
	got := setup.Init(want, nil)

	assert.Equal(t, want, got)
}

func TestInit_fails(t *testing.T) {
	assert.Panics(t, func() {
		_ = setup.Init("test", errors.New("initialization failed"))
	})
}
