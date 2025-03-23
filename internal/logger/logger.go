package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func customLevelEnconder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[CHAT] " + l.CapitalString() + "  ")
}

func NewLogger() *zap.Logger {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:   "msg",
		TimeKey:      "",
		LevelKey:     "level",
		EncodeLevel:  customLevelEnconder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)

	return zap.New(core, zap.AddCaller())
}
