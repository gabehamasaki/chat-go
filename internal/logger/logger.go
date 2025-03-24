package logger

import (
	"os"

	"github.com/gabehamasaki/chat-go/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func customLevelEnconder(appName string) func(zapcore.Level, zapcore.PrimitiveArrayEncoder) {
	return func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + appName + "] " + l.CapitalString())
	}
}

func NewLogger(cfg *config.Config) *zap.Logger {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:   "msg",
		TimeKey:      "",
		LevelKey:     "level",
		EncodeLevel:  customLevelEnconder(cfg.NAME),
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)

	return zap.New(core, zap.AddCaller())
}
