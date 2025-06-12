package log

import (
	"os"
	"strings"
	"time"

	"github.com/mattn/go-colorable"
	"github.com/mgutz/ansi"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Instance = NewLogger()

var levelColors = map[zapcore.Level]string{
	zapcore.DebugLevel:  ansi.ColorCode("magenta"),
	zapcore.InfoLevel:   ansi.ColorCode("green"),
	zapcore.WarnLevel:   ansi.ColorCode("yellow+b"),
	zapcore.ErrorLevel:  ansi.ColorCode("red+b"),
	zapcore.DPanicLevel: ansi.ColorCode("cyan+b"),
	zapcore.PanicLevel:  ansi.ColorCode("white+b+h:red"),
	zapcore.FatalLevel:  ansi.ColorCode("white+b:red"),
}

func NewLogger() *zap.Logger {
	return zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
			TimeKey:          "time",
			LevelKey:         "level",
			NameKey:          "logger",
			CallerKey:        "caller",
			MessageKey:       "msg",
			StacktraceKey:    "stacktrace",
			LineEnding:       zapcore.DefaultLineEnding,
			ConsoleSeparator: " ",
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(ansi.ColorCode("white+b") + t.Format("2006-01-02 15:04:05") + ansi.Reset)
			},
			EncodeLevel: func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
				label := l.CapitalString()
				if color, ok := levelColors[l]; ok {
					enc.AppendString("[" + color + label + ansi.Reset + "]")
				} else {
					enc.AppendString("[" + ansi.DefaultBG + label + ansi.Reset + "]")
				}
			},
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}),
		zapcore.AddSync(colorable.NewColorableStdout()),
		func() zapcore.Level {
			switch strings.ToLower(os.Getenv("LOG_LEVEL")) {
			case "debug":
				return zapcore.DebugLevel
			case "info":
				return zapcore.InfoLevel
			case "warn", "warning":
				return zapcore.WarnLevel
			case "error":
				return zapcore.ErrorLevel
			case "dpanic":
				return zapcore.DPanicLevel
			case "panic":
				return zapcore.PanicLevel
			case "fatal":
				return zapcore.FatalLevel
			default:
				return zapcore.DebugLevel
			}
		}(),
	),
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
}
