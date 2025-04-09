package cron

import liblog "tns-energo/lib/log"

type logger struct {
	log liblog.Logger
}

func (l logger) Debug(msg string, args ...any) {
	l.log.Debugf(msg, args...)
}

func (l logger) Error(msg string, args ...any) {
	l.log.Errorf(msg, args...)
}

func (l logger) Info(msg string, args ...any) {
	l.log.Debugf(msg, args...)
}

func (l logger) Warn(msg string, args ...any) {
	l.log.Debugf(msg, args...)
}
