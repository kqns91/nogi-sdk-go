package nogi

import "context"

type Logger interface {
	Debug(msg string, args ...any)
	DebugCtx(ctx context.Context, msg string, args ...any)
	Info(msg string, args ...any)
	InfoCtx(ctx context.Context, msg string, args ...any)
	Warn(msg string, args ...any)
	WarnCtx(ctx context.Context, msg string, args ...any)
	Error(msg string, args ...any)
	ErrorCtx(ctx context.Context, msg string, args ...any)
}

type nilLogger struct {
}

func newNilLogger() Logger {
	return &nilLogger{}
}

func (l *nilLogger) Debug(msg string, args ...any) {
	return
}

func (l *nilLogger) DebugCtx(ctx context.Context, msg string, args ...any) {
	return
}

func (l *nilLogger) Info(msg string, args ...any) {
	return
}

func (l *nilLogger) InfoCtx(ctx context.Context, msg string, args ...any) {
	return
}

func (l *nilLogger) Warn(msg string, args ...any) {
	return
}

func (l *nilLogger) WarnCtx(ctx context.Context, msg string, args ...any) {
	return
}

func (l *nilLogger) Error(msg string, args ...any) {
	return
}

func (l *nilLogger) ErrorCtx(ctx context.Context, msg string, args ...any) {
	return
}
