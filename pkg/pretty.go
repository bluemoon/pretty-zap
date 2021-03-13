package pkg

import (
		"fmt"
		"go.uber.org/zap"
		"os"
		"time"

		"go.uber.org/zap/buffer"
		"go.uber.org/zap/zapcore"
)

type PrettyEncoder struct {
		*zapcore.EncoderConfig
		buf *buffer.Buffer
}

func (p *PrettyEncoder) AppendDuration(duration time.Duration) {
		cur := p.buf.Len()
		p.EncodeDuration(duration, p)
		if cur == p.buf.Len() {
				p.AppendInt64(int64(duration))
		}
}

func (p *PrettyEncoder) AppendTime(t time.Time) {
		cur := p.buf.Len()
		p.EncodeTime(t, p)
		if cur == p.buf.Len() {
				p.AppendInt64(int64(t.UnixNano()))
		}
}

func (p *PrettyEncoder) AppendArray(marshaler zapcore.ArrayMarshaler) error {
		p.buf.AppendByte('[')
		err := marshaler.MarshalLogArray(p)
		p.buf.AppendByte(']')
		return err
}

func (p *PrettyEncoder) AppendObject(marshaler zapcore.ObjectMarshaler) error {
		panic("implement me")
}

func (p *PrettyEncoder) AppendReflected(value interface{}) error {
		panic("implement me")
}

func NewPrettyEncoder(encoderConfig *zapcore.EncoderConfig) *PrettyEncoder {
		bp := _bufferPool.Get()
		return &PrettyEncoder{
				EncoderConfig: encoderConfig,
				buf:           bp,
		}
}

var _ zapcore.Encoder = (*PrettyEncoder)(nil)
var _ zapcore.PrimitiveArrayEncoder = (*PrettyEncoder)(nil)
var _ zapcore.ArrayEncoder = (*PrettyEncoder)(nil)

func (p *PrettyEncoder) addKey(key string) {
		p.buf.AppendByte(' ')
		p.buf.AppendString(White)
		p.buf.AppendString(key)
		p.buf.AppendString(Clear)
		p.buf.AppendByte('=')
}

func (p *PrettyEncoder) AppendBool(b bool) {
		p.buf.AppendBool(b)
}

func (p *PrettyEncoder) AppendByteString(bytes []byte) {
		p.buf.AppendString(string(bytes))
}

func (p *PrettyEncoder) AppendComplex128(c complex128) {
		r, i := float64(real(c)), float64(imag(c))
		p.buf.AppendFloat(r, 64)
		p.buf.AppendByte('+')
		p.buf.AppendFloat(i, 64)
		p.buf.AppendByte('i')
}

func (p *PrettyEncoder) AppendComplex64(c complex64) {
		p.AppendComplex128(complex128(c))
}

func (p *PrettyEncoder) AppendFloat64(f float64) {
		p.buf.AppendFloat(f, 64)
}

func (p *PrettyEncoder) AppendFloat32(f float32) {
		p.buf.AppendFloat(float64(f), 32)
}

func (p *PrettyEncoder) AppendInt(i int) {
		p.buf.AppendInt(int64(i))
}

func (p *PrettyEncoder) AppendInt64(i int64) {
		p.buf.AppendInt(i)
}

func (p *PrettyEncoder) AppendInt32(i int32) {
		p.buf.AppendInt(int64(i))
}

func (p *PrettyEncoder) AppendInt16(i int16) {
		p.buf.AppendInt(int64(i))
}

func (p *PrettyEncoder) AppendInt8(i int8) {
		p.buf.AppendInt(int64(i))
}

func (p *PrettyEncoder) AppendString(s string) {
		p.buf.AppendString(s)
}

func (p *PrettyEncoder) AppendUint(u uint) {
		p.buf.AppendUint(uint64(u))
}

func (p *PrettyEncoder) AppendUint64(u uint64) {
		p.buf.AppendUint(u)
}

func (p *PrettyEncoder) AppendUint32(u uint32) {
		p.buf.AppendUint(uint64(u))
}

func (p *PrettyEncoder) AppendUint16(u uint16) {
		p.buf.AppendUint(uint64(u))
}

func (p *PrettyEncoder) AppendUint8(u uint8) {
		p.buf.AppendUint(uint64(u))
}

func (p *PrettyEncoder) AppendUintptr(u uintptr) {
		p.buf.AppendUint(uint64(u))
}

func (p *PrettyEncoder) AddArray(key string, marshaler zapcore.ArrayMarshaler) error {
		p.addKey(key)
		return marshaler.MarshalLogArray(p)
}

func (p *PrettyEncoder) AddObject(key string, marshaler zapcore.ObjectMarshaler) error {
		panic("implement me")
}

func (p *PrettyEncoder) AddBinary(key string, value []byte) {
		panic("implement me")
}

func (p *PrettyEncoder) AddByteString(key string, value []byte) {
		panic("implement me")
}

func (p *PrettyEncoder) AddBool(key string, value bool) {
		panic("implement me")
}

func (p *PrettyEncoder) AddComplex128(key string, value complex128) {
		panic("implement me")
}

func (p *PrettyEncoder) AddComplex64(key string, value complex64) {
		panic("implement me")
}

func (p *PrettyEncoder) AddDuration(key string, value time.Duration) {
		panic("implement me")
}

func (p *PrettyEncoder) AddFloat64(key string, value float64) {
		panic("implement me")
}

func (p *PrettyEncoder) AddFloat32(key string, value float32) {
		panic("implement me")
}

func (p *PrettyEncoder) AddInt(key string, value int) {
		panic("implement me")
}

func (p *PrettyEncoder) AddInt64(key string, value int64) {
		p.addKey(key)
		p.AppendInt64(value)
}

func (p *PrettyEncoder) AddInt32(key string, value int32) {
		panic("implement me")
}

func (p *PrettyEncoder) AddInt16(key string, value int16) {
		panic("implement me")
}

func (p *PrettyEncoder) AddInt8(key string, value int8) {
		panic("implement me")
}

func (p *PrettyEncoder) AddString(key, value string) {
		p.addKey(key)
		p.AppendString(value)
}

func (p *PrettyEncoder) AddTime(key string, value time.Time) {
		panic("implement me")
}

func (p *PrettyEncoder) AddUint(key string, value uint) {
		panic("implement me")
}

func (p *PrettyEncoder) AddUint64(key string, value uint64) {
		panic("implement me")
}

func (p *PrettyEncoder) AddUint32(key string, value uint32) {
		panic("implement me")
}

func (p *PrettyEncoder) AddUint16(key string, value uint16) {
		panic("implement me")
}

func (p *PrettyEncoder) AddUint8(key string, value uint8) {
		panic("implement me")
}

func (p *PrettyEncoder) AddUintptr(key string, value uintptr) {
		panic("implement me")
}

func (p *PrettyEncoder) AddReflected(key string, value interface{}) error {
		panic("implement me")
}

func (p *PrettyEncoder) OpenNamespace(key string) {
		panic("implement me")
}

func (p *PrettyEncoder) clone() *PrettyEncoder {
		bp := _bufferPool.Get()
		return &PrettyEncoder{
				EncoderConfig: p.EncoderConfig,
				buf:           bp,
		}
}

func (p *PrettyEncoder) Clone() zapcore.Encoder {
		return p.clone()
}

func (p *PrettyEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
		handler := p.clone()
		handler.buf.AppendByte('[')
		if entry.Level.Enabled(entry.Level) {
				handler.buf.AppendString(ColorKey(entry.Level))
				level := fmt.Sprintf("%-5s", entry.Level.String())
				handler.buf.AppendString(level)
				handler.buf.AppendString(Clear)
		}
		handler.buf.AppendByte(' ')
		handler.EncodeTime(entry.Time, handler)
		handler.buf.AppendByte(']')

		if entry.Caller.Defined && !isEmpty(p.CallerKey) {
				handler.buf.AppendByte('[')
				handler.EncodeCaller(entry.Caller, handler)
				handler.buf.AppendByte(']')
		}
		handler.buf.AppendByte(' ')
		// add message to the log
		if !isEmpty(p.MessageKey) {
				handler.buf.AppendString(entry.Message)
		}
		// add the fields
		for i := range fields {
				fields[i].AddTo(handler)
		}

		handler.buf.AppendByte('\n')
		return handler.buf, nil
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
