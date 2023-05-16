package logging_config

import (
	"github.com/denysetiawan28/go-log/src/server/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"time"
)

type LogLevel int8

// NewZapLogger
// This constructor is to initialize zap logger and create zap instance.
// You can custom the log output.
// For default this constructor use stdout and file output to generate logs
func NewZapLogger(conf config.Logs, level LogLevel, writers ...io.Writer) (logger *zap.Logger) {

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   conf.Path,
		MaxSize:    conf.MaximumLogSize,
		MaxAge:     conf.MaximumLogAge,
		MaxBackups: conf.MaximumLogBackup,
	})

	c := zapcore.AddSync(os.Stdout)

	zapWritter := make([]zapcore.WriteSyncer, 0)
	for _, writer := range writers {
		if writer == nil {
			continue
		}

		//zapWritter = append(zapWritter, zapcore.AddSync(writer))
		zapWritter = append(zapWritter, w)
	}

	if conf.File {
		zapWritter = append(zapWritter, w)
	}
	if conf.Stdout {
		zapWritter = append(zapWritter, c)
	}

	//core :=zapcore.NewTee()
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
