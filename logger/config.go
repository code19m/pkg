// Package logger provides a structured logging interface for applications.
package logger

import (
	"github.com/code19m/errx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Config defines configuration options for the logger.
type Config struct {
	// Level specifies the minimum log level to emit.
	// Valid values are: "debug", "info", "warn", "error"
	// Default is "debug".
	Level string `yaml:"level" validate:"oneof=debug info warn error" default:"debug"`

	// Encoding specifies the log format.
	// Valid values are: "json", "console"
	// Default is "json".
	Encoding string `yaml:"encoding" validate:"oneof=json console" default:"json"`
}

// getZapConfig converts the logger Config to a zap.Config.
func (c Config) getZapConfig() (*zap.Config, error) {
	zapLevel := zap.NewAtomicLevel()

	err := zapLevel.UnmarshalText([]byte(c.Level))
	if err != nil {
		return nil, errx.Wrap(err)
	}

	zapConfig := zap.Config{
		Level:            zapLevel,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		Encoding:         c.Encoding,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "file",
			TimeKey:        "time",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.RFC3339TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeName:     zapcore.FullNameEncoder,
		},
	}

	return &zapConfig, nil
}
