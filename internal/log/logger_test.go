package log_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	"goxcrap/internal/log"
)

func TestNewLogger_successWithNilWriter(t *testing.T) {
	var buf bytes.Buffer
	log.NewLogger(&buf)

	// Replaces the previous logger, so the buffer should not have logs
	log.NewLogger(nil)

	// Write a test message
	log.Info(context.Background(), "test message")

	want := ""
	got := buf.String()

	assert.Equal(t, want, got)
}

func TestLogLevels_success(t *testing.T) {
	var buf bytes.Buffer
	log.NewLogger(&buf)

	tests := []struct {
		name        string
		level       zerolog.Level
		msg         string
		expectLevel string
		expectMsg   string
	}{
		{"Trace", zerolog.TraceLevel, "trace message", `"level":"trace"`, `"message":"trace message"`},
		{"Debug", zerolog.DebugLevel, "debug message", `"level":"debug"`, `"message":"debug message"`},
		{"Info", zerolog.InfoLevel, "info message", `"level":"info"`, `"message":"info message"`},
		{"Warn", zerolog.WarnLevel, "warn message", `"level":"warn"`, `"message":"warn message"`},
		{"Error", zerolog.ErrorLevel, "error message", `"level":"error"`, `"message":"error message"`},
		{"Fatal", zerolog.FatalLevel, "fatal message", `"level":"fatal"`, `"message":"fatal message"`},
		{"Panic", zerolog.PanicLevel, "panic message", `"level":"panic"`, `"message":"panic message"`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf.Reset()

			// Call the appropriate function based on the level
			switch tt.level {
			case zerolog.TraceLevel:
				log.Trace(context.Background(), tt.msg)
			case zerolog.DebugLevel:
				log.Debug(context.Background(), tt.msg)
			case zerolog.InfoLevel:
				log.Info(context.Background(), tt.msg)
			case zerolog.WarnLevel:
				log.Warn(context.Background(), tt.msg)
			case zerolog.ErrorLevel:
				log.Error(context.Background(), tt.msg)
			case zerolog.FatalLevel:
				log.Fatal(context.Background(), tt.msg)
			case zerolog.PanicLevel:
				log.Panic(context.Background(), tt.msg)
			default:
				t.Error("Wrong log level")
			}

			assert.Contains(t, buf.String(), tt.expectLevel)
			assert.Contains(t, buf.String(), tt.expectMsg)
		})
	}
}

func TestErr_success(t *testing.T) {
	var buf bytes.Buffer
	log.NewLogger(&buf)

	tests := []struct {
		name           string
		err            error
		msg            string
		expectLevel    string
		expectMsg      string
		expectErrorMsg string
	}{
		{"ErrorWithErr", assert.AnError, "error message", `"level":"error"`, `"message":"error message"`, `"error":"assert.AnError"`},
		{"ErrorWithoutErr", nil, "info message", `"level":"info"`, `"message":"info message"`, ``},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf.Reset()

			log.Err(context.Background(), tt.err, tt.msg)

			assert.Contains(t, buf.String(), tt.expectLevel)
			assert.Contains(t, buf.String(), tt.expectMsg)
		})
	}
}
