package log

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Instance = ZapLogger()

func ZapLogger() *zap.Logger {
	encoder := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString("\x1b[90m" + t.Format("2006-01-02 15:04:05") + "\x1b[0m")
		},
		EncodeLevel: func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			var color string
			switch l {
			case zapcore.DebugLevel:
				color = "\x1b[35m"
			case zapcore.InfoLevel:
				color = "\x1b[32m"
			case zapcore.WarnLevel:
				color = "\x1b[33m"
			case zapcore.ErrorLevel:
				color = "\x1b[31m"
			default:
				color = "\x1b[0m"
			}
			enc.AppendString(color + "[" + l.CapitalString() + "]" + "\x1b[0m")
		},
		ConsoleSeparator: " ",
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoder),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)

	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}
