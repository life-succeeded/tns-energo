package log

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	WithSkip(s int) Logger

	// WithUserInfo Возвращает Logger с сохраненными полями userId и email.
	// Эти поля будут попадать в каждую запись, сделанную полученным логгером
	WithUserInfo(userId int, email string) Logger

	// WithTags Возвращает Logger c новыми тегами, имеющиеся теги из оригинального Logger будут скопированы в новый
	WithTags(tags ...Tag) Logger

	// Message возвращает LogEntry ассоциированный с указанным уровнем
	Message(level MessageLevel, message string) LogEntry

	// DebugEntry возвращает LogEntry ассоциированный с указанным сообщением и уровнем лога LevelDebug
	DebugEntry(message string) LogEntry
	// DebugEntryf форматирует текст и возвращает LogEntry ассоциированный с указанным сообщением и уровнем лога LevelDebug
	DebugEntryf(format string, args ...interface{}) LogEntry

	// Debug записывает сообщение в лог в момент вызова
	Debug(message string)
	// Debugf форматирует сообщение и записывает его в момент вызова
	Debugf(format string, args ...interface{})

	// ErrorEntry возвращает LogEntry ассоциированный с указанным сообщением и уровнем лога LevelError
	ErrorEntry(message string) LogEntry
	// ErrorEntryf форматирует текст и возвращает LogEntry ассоциированный с указанным сообщением и уровнем лога LevelError
	ErrorEntryf(format string, args ...interface{}) LogEntry

	// Error записывает сообщение в лог в момент вызова
	Error(message string)
	// Errorf форматирует сообщение и записывает его в момент вызова
	Errorf(format string, args ...interface{})
}

var (
	_global Logger
)

func Global() Logger {
	if _global == nil {
		_global = NewLogger("default")
	}

	return _global
}

// LoggerImpl реализация Logger основанная на zap.Logger
type LoggerImpl struct {
	log    *zap.Logger
	nowLog *zap.Logger
	tags   []Tag
}

// NewLogger возвращает новый экземпляр логгера
func NewLogger(appName string, options ...Option) LoggerImpl {
	jsonConfig := zapcore.EncoderConfig{
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeTime:    zapcore.RFC3339TimeEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
		MessageKey:    "message",
		LevelKey:      "level",
		TimeKey:       "date",
		CallerKey:     "caller",
		StacktraceKey: "stackTrace",
	}

	optionsHolder := optionHolder{
		out: os.Stdout,
		err: os.Stderr,
	}
	for _, option := range options {
		optionsHolder = option.apply(optionsHolder)
	}

	skip := 1 // log.Debugf -> log.writeEntry
	if optionsHolder.skip > 0 {
		skip += optionsHolder.skip
	}

	zapOptions := []zap.Option{
		zap.ErrorOutput(newWriteSyncer(optionsHolder.err)),
		zap.AddCaller(),
		zap.AddCallerSkip(skip),
	}

	if optionsHolder.stacktrace {
		zapOptions = append(zapOptions, zap.AddStacktrace(_swStacktraceEnabler))
	}
	containerId, _ := os.Hostname()

	log := zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(jsonConfig),
			newWriteSyncer(optionsHolder.out),
			_swLevelEnabler),
		zapOptions...,
	).With(
		zap.String("appName", appName),
		zap.String("containerId", containerId),
	)

	logger := LoggerImpl{
		log:    log.WithOptions(zap.AddCallerSkip(1)),
		nowLog: log,
		tags:   optionsHolder.tags,
	}

	if optionsHolder.global {
		_global = logger
	}

	return logger
}

func (s LoggerImpl) WithSkip(n int) Logger {
	s.nowLog = s.nowLog.WithOptions(zap.AddCallerSkip(n))
	s.log = s.nowLog.WithOptions(zap.AddCallerSkip(1))
	return s
}

func (s LoggerImpl) WithUserInfo(userId int, email string) Logger {
	if userId <= 0 || len(email) == 0 {
		return s
	}

	fields := []zap.Field{
		zap.Int("userId", userId),
		zap.String("email", email),
	}
	s.log = s.log.With(fields...)
	s.nowLog = s.nowLog.With(fields...)

	return s
}

func (s LoggerImpl) WithTags(tags ...Tag) Logger {
	s.tags = append(s.tags, tags...)
	return s
}

func (s LoggerImpl) Message(level MessageLevel, message string) LogEntry {
	return Entry{
		level:     level,
		message:   message,
		writeFunc: s.writeEntry,
	}
}

func (s LoggerImpl) DebugEntry(message string) LogEntry {
	return Entry{
		level:     LevelDebug,
		message:   message,
		tags:      s.tags,
		writeFunc: s.writeEntry,
	}
}

func (s LoggerImpl) DebugEntryf(format string, args ...interface{}) LogEntry {
	return s.DebugEntry(fmt.Sprintf(format, args...))
}

func (s LoggerImpl) Debug(message string) {
	s.nowLog.Debug(message, marshalTags(s.tags))
}

func (s LoggerImpl) Debugf(format string, args ...interface{}) {
	s.nowLog.Debug(fmt.Sprintf(format, args...), marshalTags(s.tags))
}

func (s LoggerImpl) ErrorEntry(message string) LogEntry {
	return Entry{
		level:     LevelError,
		message:   message,
		tags:      s.tags,
		writeFunc: s.writeEntry,
	}
}

func (s LoggerImpl) ErrorEntryf(format string, args ...interface{}) LogEntry {
	return s.ErrorEntry(fmt.Sprintf(format, args...))
}

func (s LoggerImpl) Error(message string) {
	s.nowLog.Error(message, marshalTags(s.tags))
}

func (s LoggerImpl) Errorf(format string, args ...interface{}) {
	s.nowLog.Error(fmt.Sprintf(format, args...), marshalTags(s.tags))
}

func (s LoggerImpl) writeEntry(entry Entry) {
	switch entry.level {
	case LevelDebug:
		s.log.Debug(entry.message, marshalTags(entry.tags))
	case LevelError:
		s.log.Error(entry.message, marshalTags(entry.tags))
	}
}

func marshalTags(tags []Tag) zap.Field {
	return zap.Array("tags", zapcore.ArrayMarshalerFunc(func(e zapcore.ArrayEncoder) error {
		for _, tag := range tags {
			if len(tag) == 0 {
				continue
			}

			e.AppendString(string(tag))
		}

		return nil
	}))
}
