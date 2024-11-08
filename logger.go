package main

import "github.com/bitrise-io/go-utils/v2/log"

type deviceFinderLogger struct {
	logger log.Logger
}

func newDeviceFinderLogger(logger log.Logger) log.Logger {
	return deviceFinderLogger{
		logger: logger,
	}
}

func (l deviceFinderLogger) Infof(format string, v ...interface{}) {
	l.logger.Debugf(format, v...)
}

func (l deviceFinderLogger) Warnf(format string, v ...interface{}) {
	l.logger.Debugf(format, v...)
}

func (l deviceFinderLogger) Printf(format string, v ...interface{}) {
	l.logger.Debugf(format, v...)
}

func (l deviceFinderLogger) Donef(format string, v ...interface{}) {
	l.logger.Debugf(format, v...)
}

func (l deviceFinderLogger) Debugf(format string, v ...interface{}) {
	l.logger.Debugf(format, v...)
}

func (l deviceFinderLogger) Errorf(format string, v ...interface{}) {
	l.logger.Debugf(format, v...)
}

func (l deviceFinderLogger) TInfof(format string, v ...interface{}) {
	l.logger.TDebugf(format, v...)
}

func (l deviceFinderLogger) TWarnf(format string, v ...interface{}) {
	l.logger.TDebugf(format, v...)
}

func (l deviceFinderLogger) TPrintf(format string, v ...interface{}) {
	l.logger.TDebugf(format, v...)
}

func (l deviceFinderLogger) TDonef(format string, v ...interface{}) {
	l.logger.TDebugf(format, v...)
}

func (l deviceFinderLogger) TDebugf(format string, v ...interface{}) {
	l.logger.TDebugf(format, v...)
}

func (l deviceFinderLogger) TErrorf(format string, v ...interface{}) {
	l.logger.TDebugf(format, v...)
}

func (l deviceFinderLogger) Println() {
	l.logger.Debugf("")
}

func (l deviceFinderLogger) EnableDebugLog(enable bool) {
	l.logger.EnableDebugLog(enable)
}
