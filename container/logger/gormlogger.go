package logger

import (
	"context"
	"fmt"
	gormLogger "gorm.io/gorm/logger"
	gormUtils "gorm.io/gorm/utils"
	"time"
)

// Customize SQL Logger for gorm library
// ref: https://github.com/wantedly/gorm-zap
// ref: https://github.com/go-gorm/gorm/blob/master/logger/logger.go

const (
	logTitle      = "[gorm] "
	sqlFormat     = logTitle + "%s"
	messageFormat = logTitle + "%s, %s"
	errorFormat   = logTitle + "%s, %s, %s"
	slowThreshold = 200
)

// LogMode The log level of gorm logger is overwritten by the log level of Zap logger
func (l *logger) LogMode(_ gormLogger.LogLevel) gormLogger.Interface {
	return l
}

// Info prints information log
func (l *logger) Info(_ context.Context, s string, data ...any) {
	l.ZapLogger.Infof(messageFormat, append([]any{s, gormUtils.FileWithLineNum()}, data...)...)
}

// Warn prints a warning log.
func (l *logger) Warn(_ context.Context, msg string, data ...interface{}) {
	l.ZapLogger.Warnf(messageFormat, append([]interface{}{msg, gormUtils.FileWithLineNum()}, data...)...)
}

// Error prints a error log.
func (l *logger) Error(_ context.Context, msg string, data ...interface{}) {
	l.ZapLogger.Errorf(messageFormat, append([]interface{}{msg, gormUtils.FileWithLineNum()}, data...)...)
}

// Trace prints a trace log such as sql, source file and error.
func (l *logger) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)

	switch {
	case err != nil:
		sql, _ := fc()
		l.GetZapLogger().Errorf(errorFormat, gormUtils.FileWithLineNum(), err, sql)
	case elapsed > slowThreshold*time.Millisecond:
		sql, _ := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", slowThreshold)
		l.GetZapLogger().Warnf(errorFormat, gormUtils.FileWithLineNum(), slowLog, sql)
	default:
		sql, _ := fc()
		l.GetZapLogger().Debugf(sqlFormat, sql)
	}
}
