package log_test

import (
	"context"
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

func TestRetrieveParam_success(t *testing.T) {
	ctx := context.Background()
	ctx = log.With(ctx, log.Param("test_param", "test"))

	want := "test"
	got := log.RetrieveParam[string](ctx, "test_param")

	assert.Equal(t, want, got)
}

func TestRetrieveParam_successWithZeroValueWhenTheParamsMapHasNotBeenCreated(t *testing.T) {
	want := ""
	got := log.RetrieveParam[string](context.Background(), "test_param")

	assert.Equal(t, want, got)
}

func TestRetrieveParam_successWithZeroValueWhenTheKeyIsNotPresent(t *testing.T) {
	ctx := log.With(context.Background(), log.Param("test_param", "test"))

	want := 0
	got := log.RetrieveParam[int](ctx, "nonexistent_key")

	assert.Equal(t, want, got)
}

func TestRetrieveParam_successWithZeroValueWhenTheKeyIsNotTheExpectedType(t *testing.T) {
	ctx := log.With(context.Background(), log.Param("test_param", "test"))

	want := 0
	got := log.RetrieveParam[int](ctx, "test_param")

	assert.Equal(t, want, got)
}
