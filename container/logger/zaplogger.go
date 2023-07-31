package logger

import (
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

func newZapLogger(cfg *Config) (*zap.Logger, error) {
	zapCfg := cfg.ZapConfig
	enc, _ := newEncoder(&zapCfg)
	writer := open(zapCfg.OutputPaths, &cfg.LogRotate)

	if zapCfg.Level == (zap.AtomicLevel{}) {
		return nil, errors.New("missing Level")
	}

	log := zap.New(zapcore.NewCore(enc, writer, zapCfg.Level), zap.AddCaller(), zap.AddCallerSkip(1))
	return log, nil
}

func newEncoder(cfg *zap.Config) (zapcore.Encoder, error) {
	// defaultConfig uses the basic production environment configuration provided by zap
	defaultConfig := zap.NewProductionEncoderConfig()
	// Time format
	defaultConfig.EncodeTime = zapLogLocalTimeEncoder
	// Level capital letters
	defaultConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	switch cfg.Encoding {
	case "console":
		return zapcore.NewConsoleEncoder(defaultConfig), nil
	case "json":
		return zapcore.NewJSONEncoder(defaultConfig), nil
	}
	return nil, errors.New("failed to set encoder")
}

func open(paths []string, rotateCfg *lumberjack.Logger) zapcore.WriteSyncer {
	writers := make([]zapcore.WriteSyncer, 0, len(paths))
	for _, path := range paths {
		writer := newWriter(path, rotateCfg)
		writers = append(writers, writer)
	}
	writer := zap.CombineWriteSyncers(writers...)
	return writer
}

func newWriter(path string, rotateCfg *lumberjack.Logger) zapcore.WriteSyncer {
	switch path {
	case "stdout":
		return os.Stdout
	case "stderr":
		return os.Stderr
	}
	sink := zapcore.AddSync(
		&lumberjack.Logger{
			Filename:   path,
			MaxSize:    rotateCfg.MaxSize,
			MaxBackups: rotateCfg.MaxBackups,
			MaxAge:     rotateCfg.MaxAge,
			Compress:   rotateCfg.Compress,
		},
	)
	return sink
}

const MsecLocalTimeFormat = "2006-01-02 15:04:05.000"

// zapLogLocalTimeEncoder custom local time format
func zapLogLocalTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(MsecLocalTimeFormat))
}
