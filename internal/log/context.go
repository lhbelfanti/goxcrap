package log

import (
	"context"
	"time"

	"github.com/rs/zerolog"
)

// field represent a key-value tuple that will be added to the context
type (
	field struct {
		Key   string
		Value interface{}
	}

	logCtxKey struct{}
)

// Param creates a new field to be saved into context
func Param(key string, value interface{}) field {
	return field{key, value}
}

// With custom function to add log parameters to the context
func With(ctx context.Context, fields ...field) context.Context {
	// Get the existing map of parameters or create a new one
	params, ok := ctx.Value(logCtxKey{}).(map[string]interface{})
	if !ok {
		params = make(map[string]interface{}, len(fields))
	}

	for _, t := range fields {
		params[t.Key] = t.Value
	}

	return context.WithValue(ctx, logCtxKey{}, params)
}

func withContextParams(ctx context.Context, event *zerolog.Event) *zerolog.Event {
	params, ok := ctx.Value(logCtxKey{}).(map[string]interface{})
	if !ok {
		return event
	}

	// If a specific type (supported by zerolog.Event) is needed, add it to the switch
	for key, value := range params {
		switch v := value.(type) {
		case string:
			event = event.Str(key, v)
		case int:
			event = event.Int(key, v)
		case float64:
			event = event.Float64(key, v)
		case bool:
			event = event.Bool(key, v)
		case error:
			event = event.AnErr(key, v)
		case []string:
			event = event.Strs(key, v)
		case []int:
			event = event.Ints(key, v)
		case []float64:
			event = event.Floats64(key, v)
		case []byte:
			event = event.Bytes(key, v)
		case time.Time:
			event = event.Time(key, v)
		default:
			event = event.Interface(key, v)
		}
	}

	return event
}
