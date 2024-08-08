package log

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWithContextParams_success(t *testing.T) {
	tests := []struct {
		name string
		ctx  context.Context
		want map[string]interface{}
	}{
		{
			name: "Empty context",
			ctx:  context.Background(),
			want: map[string]interface{}{},
		},
		{
			name: "Single types",
			ctx: With(context.Background(),
				Param("stringKey", "stringValue"),
				Param("intKey", 42),
				Param("float64Key", 42.5),
				Param("boolKey", true),
			),
			want: map[string]interface{}{
				"stringKey":  "stringValue",
				"intKey":     42,
				"float64Key": 42.5,
				"boolKey":    true,
			},
		},
		{
			name: "Error value",
			ctx:  With(context.Background(), Param("errorKey", assert.AnError)),
			want: map[string]interface{}{"errorKey": assert.AnError},
		},
		{
			name: "Array values",
			ctx: With(context.Background(),
				Param("stringArrayKey", []string{"value1", "value2"}),
				Param("intArrayKey", []int{1, 2, 3}),
				Param("float64ArrayKey", []float64{1.0, 2.0, 3.0}),
				Param("bytesKey", []byte("value")),
			),
			want: map[string]interface{}{
				"stringArrayKey":  []string{"value1", "value2"},
				"intArrayKey":     []int{1, 2, 3},
				"float64ArrayKey": []float64{1.0, 2.0, 3.0},
				"bytesKey":        []byte("value"),
			},
		},
		{
			name: "Time value",
			ctx:  With(context.Background(), Param("timeKey", time.Now())),
			want: map[string]interface{}{"timeKey": time.Now()},
		},
		{
			name: "Interface value",
			ctx: With(context.Background(),
				Param("interfaceKey", struct{ key string }{key: "value"}),
			),
			want: map[string]interface{}{
				"interfaceKey": struct{ key string }{key: "value"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			NewLogger(&buf)
			Info(tt.ctx, "")

			got := buf.String()

			for key, _ := range tt.want {
				assert.Contains(t, got, key)
			}
		})
	}
}
func TestWith_success(t *testing.T) {
	ctx := context.Background()
	want := []struct {
		Key   string
		Value interface{}
	}{
		{"key1", "value1"},
		{"key2", 123},
		{"key1", "newValue"},
	}

	// Empty context
	field1 := Param(want[0].Key, want[0].Value)
	field2 := Param(want[1].Key, want[1].Value)

	ctx = With(ctx, field1, field2)

	got, ok := ctx.Value(logCtxKey{}).(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, want[0].Value, got[want[0].Key])
	assert.Equal(t, want[1].Value, got[want[1].Key])

	// Context with params added
	field3 := Param(want[2].Key, want[2].Value)
	ctx = With(ctx, field3)

	got, ok = ctx.Value(logCtxKey{}).(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, want[2].Value, got[want[2].Key])
}
