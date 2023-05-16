package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/net/context"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"time"
)

type Level int8

type defaultLogger struct {
	writers   []io.Writer
	zapLogger *zap.Logger
	level     Level
}

type combineLogger struct {
	sysLog Logger
	//tdrLog Logger
}

type Field struct {
	Key string
	Val interface{}
}

type Logger interface {
	Debug(ctx context.Context, message string, fields ...Field)
	Info(ctx context.Context, message string, fields ...Field)
	Warn(ctx context.Context, message string, fields ...Field)
	Error(ctx context.Context, message string, fields ...Field)
	Fatal(ctx context.Context, message string, fields ...Field)
	Panic(ctx context.Context, message string, fields ...Field)
	//TDR(ctx context.Context, tdr LogTdrModel)
	Close() error
}

func newLogger() (*zap.Logger, error) {
	defLogger := &defaultLogger{
		writers: make([]io.Writer, 0),
	}

	defLogger.zapLogger = NewZapLogger(defLogger.level, defLogger.writers...)

	return defLogger.zapLogger, nil
}

func SetupLogger() *zap.Logger {
	sysLog, _ := newLogger()

	return sysLog
}

func NewZapLogger(level Level, writers ...io.Writer) (logger *zap.Logger) {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "C:\\logs\\test.log",
		MaxSize:    10,
		MaxAge:     30,
		MaxBackups: 3,
	})

	zapWritter := make([]zapcore.WriteSyncer, 0)
	for _, writer := range writers {
		if writer == nil {
			continue
		}

		//zapWritter = append(zapWritter, zapcore.AddSync(writer))
		zapWritter = append(zapWritter, w)
	}
	zapWritter = append(zapWritter, w)

	core := zapcore.NewCore(
		getEncoder(),
		zapcore.NewMultiWriteSyncer(zapWritter...),
		zapcore.InfoLevel,
	)

	logger = zap.New(core)
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "xtime",
		MessageKey:     "x",
		EncodeDuration: millisDurationEncoder,
		EncodeTime:     timeEncoder,
		LineEnding:     zapcore.DefaultLineEnding,
	}

	return zapcore.NewJSONEncoder(encoderConfig)
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.999"))
}

func millisDurationEncoder(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt64(d.Nanoseconds() / 1000000)
}
