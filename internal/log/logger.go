package log

import (
	"context"
	"io"
	"os"

	"github.com/rs/zerolog"
)

var logger = zerolog.New(os.Stdout).With().Timestamp().Logger().Level(zerolog.DebugLevel)

// NewCustomLogger returns a new custom logger writing to the provided writer
func NewCustomLogger(writer io.Writer, logLevel zerolog.Level) {
	if writer == nil {
		writer = os.Stdout
	}

	logger = zerolog.New(writer).With().Timestamp().Logger().Level(logLevel)
}

// Trace starts a new message with trace level
func Trace(ctx context.Context, msg string) {
	logWithLevel(ctx, zerolog.TraceLevel, msg)
}

// Debug starts a new message with debug level
func Debug(ctx context.Context, msg string) {
	logWithLevel(ctx, zerolog.DebugLevel, msg)
}

// Info starts a new message with info level
func Info(ctx context.Context, msg string) {
	logWithLevel(ctx, zerolog.InfoLevel, msg)
}

// Warn starts a new message with warn level
func Warn(ctx context.Context, msg string) {
	logWithLevel(ctx, zerolog.WarnLevel, msg)
}

// Err starts a new message with error level with err as a field if not nil or
// with info level if err is nil
func Err(ctx context.Context, err error, msg string) {
	event := withContextParams(ctx, logger.Err(err))
	event.Msg(msg)
}

// Error starts a new message with error level
func Error(ctx context.Context, msg string) {
	logWithLevel(ctx, zerolog.ErrorLevel, msg)
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method
func Fatal(ctx context.Context, msg string) {
	logWithLevel(ctx, zerolog.FatalLevel, msg)
}

// Panic starts a new message with panic level. The message is also sent
// to the panic function
func Panic(ctx context.Context, msg string) {
	logWithLevel(ctx, zerolog.PanicLevel, msg)
}

// logWithLevel is a generic function to log messages at different levels
func logWithLevel(ctx context.Context, level zerolog.Level, msg string) {
	event := logger.WithLevel(level)
	event = withContextParams(ctx, event)
	event.Msg(msg)
}
