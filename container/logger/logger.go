package logger

import (
	"context"
	"fmt"
	"github.com/devops-codegpt/server/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

type Config struct {
	ZapConfig zap.Config
	LogRotate lumberjack.Logger
}

// Logger is an zap logger, it's an alternative implementation of *gorm.Logger
type Logger interface {
	GetZapLogger() *zap.SugaredLogger
	LogMode(level gormLogger.LogLevel) gormLogger.Interface
	Info(ctx context.Context, msg string, data ...interface{})
	Warn(ctx context.Context, msg string, data ...interface{})
	Error(ctx context.Context, msg string, data ...interface{})
	Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error)
}

type logger struct {
	ZapLogger *zap.SugaredLogger
}

func (l *logger) GetZapLogger() *zap.SugaredLogger {
	return l.ZapLogger
}

func InitLogger(cfg *config.Configuration) (Logger, error) {
	logCfg := cfg.Logs
	// Add multi output log paths
	filepath := getLogFilename(logCfg.Path)
	outPaths := []string{"stdout", filepath}

	// New Config
	myConfig := &Config{
		ZapConfig: zap.Config{
			Level:         zap.NewAtomicLevelAt(logCfg.Level),
			Encoding:      logCfg.Encoding,
			EncoderConfig: zapcore.EncoderConfig{},
			OutputPaths:   outPaths,
		},
		LogRotate: lumberjack.Logger{
			MaxBackups: logCfg.MaxBackups,
			MaxAge:     logCfg.MaxAge,
			Compress:   logCfg.Compress,
			MaxSize:    logCfg.MaxSize,
		},
	}
	// Get zap logger
	zapLog, err := newZapLogger(myConfig)
	if err != nil {
		fmt.Printf("Failed to compose zap logger : %s", err)
		return nil, err
	}

	sugar := zapLog.Sugar()
	// Set package variable logger
	log := NewLogger(sugar)
	log.GetZapLogger().Infof("Success to read zap logger configuration")
	_ = zapLog.Sync()
	return log, nil
}

// NewLogger is constructor for logger
func NewLogger(sugar *zap.SugaredLogger) Logger {
	return &logger{ZapLogger: sugar}
}

func getLogFilename(dir string) string {
	now := time.Now()
	return fmt.Sprintf("%s/%04d-%02d-%02d.log", dir, now.Year(), now.Month(), now.Day())
}
