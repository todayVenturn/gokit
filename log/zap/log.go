package zap

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level = zapcore.Level

const (
	DebugLevel = zapcore.DebugLevel
	InfoLevel  = zapcore.InfoLevel
	WarnLevel  = zapcore.WarnLevel
	ErrorLevel = zapcore.ErrorLevel
	PanicLevel = zapcore.PanicLevel
	FatalLevel = zapcore.FatalLevel
)

type Logger struct {
	l *zap.Logger
	s *zap.SugaredLogger
	// https://pkg.go.dev/go.uber.org/zap#example-AtomicLevel
	al *zap.AtomicLevel
}

func New(out io.Writer, level Level, opts ...Option) *Logger {
	if out == nil {
		out = os.Stderr
	}

	al := zap.NewAtomicLevelAt(level)
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.FunctionKey = "F"

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg),
		zapcore.AddSync(out),
		al,
	)
	return &Logger{l: zap.New(core, opts...), s: zap.New(core, opts...).Sugar(), al: &al}
}

// SetLevel 动态更改日志级别
// 对于使用 NewTee 创建的 Logger 无效，因为 NewTee 本意是根据不同日志级别
// 创建的多个 zap.Core，不应该通过 SetLevel 将多个 zap.Core 日志级别统一
func (l *Logger) SetLevel(level Level) {
	if l.al != nil {
		l.al.SetLevel(level)
	}
}

type Field = zap.Field

func (l *Logger) L() *zap.Logger {
	return l.l
}

func (l *Logger) S() *zap.SugaredLogger {
	return l.s
}

func (l *Logger) Debug(args ...interface{}) {
	l.s.Debugln(args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.s.Infoln(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.s.Warnln(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.s.Errorln(args...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.s.Panicln(args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.s.Fatalln(args...)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.s.Debugf(template, args...)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.s.Infof(template, args...)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.s.Warnf(template, args...)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.s.Errorf(template, args...)
}

func (l *Logger) Panicf(template string, args ...interface{}) {
	l.s.Debugf(template, args...)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.s.Fatalf(template, args...)
}

func (l *Logger) Sync() error {
	return l.s.Sync()
}

var std = New(os.Stderr, InfoLevel)

func Default() *Logger         { return std }
func ReplaceDefault(l *Logger) { std = l }

func SetLevel(level Level) { std.SetLevel(level) }

func Debug(args ...interface{}) { std.Debug(args...) }
func Info(args ...interface{})  { std.Debug(args...) }
func Warn(args ...interface{})  { std.Debug(args...) }
func Error(args ...interface{}) { std.Debug(args...) }
func Panic(args ...interface{}) { std.Debug(args...) }
func Fatal(args ...interface{}) { std.Debug(args...) }

func Debugf(template string, args ...interface{}) { std.Debugf(template, args...) }
func Infof(template string, args ...interface{})  { std.Infof(template, args...) }
func Warnf(template string, args ...interface{})  { std.Warnf(template, args...) }
func Errorf(template string, args ...interface{}) { std.Errorf(template, args...) }
func Panicf(template string, args ...interface{}) { std.Panicf(template, args...) }
func Fatalf(template string, args ...interface{}) { std.Fatalf(template, args...) }

func Sync() error { return std.Sync() }
