package log

import (
	"io"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type swWriteSyncer struct {
	w  io.Writer
	mx sync.RWMutex
}

func newWriteSyncer(w io.Writer) zapcore.WriteSyncer {
	return &swWriteSyncer{
		w: w,
	}
}

func (s *swWriteSyncer) Write(p []byte) (int, error) {
	s.mx.Lock()
	defer s.mx.Unlock()
	return s.w.Write(p)
}

func (s *swWriteSyncer) Sync() error {
	return nil
}

var (
	_swLevelEnabler = zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l == zapcore.DebugLevel || l == zapcore.ErrorLevel
	})
	_swStacktraceEnabler = zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l == zapcore.ErrorLevel
	})
)
