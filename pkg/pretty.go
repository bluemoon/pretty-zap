package pkg

import (
		"go.uber.org/zap"
		"go.uber.org/zap/buffer"
		"go.uber.org/zap/zapcore"
		"os"
)

type PrettyEncoder struct {
		*zapcore.EncoderConfig
		*Encoder
		buf *buffer.Buffer
}

func NewPrettyEncoder(encoderConfig *zapcore.EncoderConfig) *PrettyEncoder {
		bp := _bufferPool.Get()
		return &PrettyEncoder{
				EncoderConfig: encoderConfig,
				Encoder: &Encoder{
						EncoderConfig: encoderConfig,
						buf:           bp,
				},
				buf: bp,
		}
}

func (p *PrettyEncoder) Logger(level zapcore.Level) *zap.Logger {
		pe := zap.NewDevelopmentEncoderConfig()
		pe.EncodeCaller = zapcore.ShortCallerEncoder
		pe.EncodeTime = zapcore.ISO8601TimeEncoder
		pe.CallerKey = "caller"

		consoleEncoder := NewPrettyEncoder(&pe)
		zc := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level)
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
