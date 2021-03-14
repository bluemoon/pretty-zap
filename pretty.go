package prettyzap

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

type PrettyLogger struct {
	*Encoder
}

func NewPrettyLogger(encoderConfig *zapcore.EncoderConfig, level zapcore.Level) *zap.Logger {
	bp := _bufferPool.Get()

	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.CallerKey = "caller"

	pe := &PrettyLogger{
		Encoder: &Encoder{
			EncoderConfig: encoderConfig,
			buf:           bp,
		},
	}

	zc := zapcore.NewCore(pe, zapcore.AddSync(os.Stdout), level)
	l := zap.New(zc,
		zap.AddCaller(),
		zap.AddStacktrace(zap.PanicLevel),
	)
	return l
}

// thanks i hate it
var (
	_bufferPool = buffer.NewPool()
)
