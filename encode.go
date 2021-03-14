package prettyzap

import (
	"encoding/base64"
	"fmt"
	"time"

	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

func (p *Encoder) addKey(key string) {
	p.buf.AppendByte(' ')
	p.buf.AppendString(White)
	p.buf.AppendString(key)
	p.buf.AppendString(Clear)
	p.buf.AppendByte('=')
}

func (p *Encoder) addElementSeparator() {
	last := p.buf.Len() - 1
	if last < 0 {
		return
	}
	switch p.buf.Bytes()[last] {
	case '{', '[', ':', ',', ' ', '=':
		return
	default:
		p.buf.AppendByte(',')
	}
}

func (p *Encoder) c128(c complex128) {
	r, i := float64(real(c)), float64(imag(c))
	p.buf.AppendFloat(r, 64)
	p.buf.AppendByte('+')
	p.buf.AppendFloat(i, 64)
	p.buf.AppendByte('i')
}

type Encoder struct {
	*zapcore.EncoderConfig
	buf *buffer.Buffer
}

var _ zapcore.PrimitiveArrayEncoder = (*Encoder)(nil)
var _ zapcore.ArrayEncoder = (*Encoder)(nil)
var _ zapcore.Encoder = (*Encoder)(nil)

func (p *Encoder) AppendDuration(duration time.Duration) {
	p.addElementSeparator()
	cur := p.buf.Len()
	p.EncodeDuration(duration, p)
	if cur == p.buf.Len() {
		p.AppendInt64(int64(duration))
	}
}

func (p *Encoder) AppendTime(t time.Time) {
	p.addElementSeparator()
	cur := p.buf.Len()
	p.EncodeTime(t, p)
	if cur == p.buf.Len() {
		p.AppendInt64(int64(t.UnixNano()))
	}
}

func (p *Encoder) AppendArray(marshaler zapcore.ArrayMarshaler) error {
	p.addElementSeparator()
	p.buf.AppendByte('[')
	err := marshaler.MarshalLogArray(p)
	p.buf.AppendByte(']')
	return err
}

func (p *Encoder) AppendObject(marshaler zapcore.ObjectMarshaler) error {
	p.addElementSeparator()
	p.buf.AppendByte('{')
	err := marshaler.MarshalLogObject(p)
	p.buf.AppendByte('}')
	return err
}

func (p *Encoder) AppendReflected(value interface{}) error {
	p.addElementSeparator()
	v, ok := value.(string)
	if !ok {
		v = fmt.Sprintf("%v", value)
	}
	p.AppendString(v)
	return nil
}

func (p *Encoder) AppendBool(b bool) {
	p.addElementSeparator()
	p.buf.AppendBool(b)
}

func (p *Encoder) AppendByteString(bytes []byte) {
	p.addElementSeparator()
	p.buf.AppendString(string(bytes))
}

func (p *Encoder) AppendComplex128(c complex128) {
	p.addElementSeparator()
	c128(p.buf, c)
}

func (p *Encoder) AppendComplex64(c complex64) {
	p.addElementSeparator()
	c128(p.buf, complex128(c))
}

func (p *Encoder) AppendFloat64(f float64) {
	p.addElementSeparator()
	p.buf.AppendFloat(f, 64)
}

func (p *Encoder) AppendFloat32(f float32) {
	p.addElementSeparator()
	p.buf.AppendFloat(float64(f), 32)
}

func (p *Encoder) AppendInt(i int) {
	p.addElementSeparator()
	p.buf.AppendInt(int64(i))
}

func (p *Encoder) AppendInt64(i int64) {
	p.addElementSeparator()
	p.buf.AppendInt(i)
}

func (p *Encoder) AppendInt32(i int32) {
	p.addElementSeparator()
	p.buf.AppendInt(int64(i))
}

func (p *Encoder) AppendInt16(i int16) {
	p.addElementSeparator()
	p.buf.AppendInt(int64(i))
}

func (p *Encoder) AppendInt8(i int8) {
	p.addElementSeparator()
	p.buf.AppendInt(int64(i))
}

func (p *Encoder) AppendString(s string) {
	p.addElementSeparator()
	p.buf.AppendString(s)
}

func (p *Encoder) AppendUint(u uint) {
	p.addElementSeparator()
	p.buf.AppendUint(uint64(u))
}

func (p *Encoder) AppendUint64(u uint64) {
	p.addElementSeparator()
	p.buf.AppendUint(u)
}

func (p *Encoder) AppendUint32(u uint32) {
	p.addElementSeparator()
	p.buf.AppendUint(uint64(u))
}

func (p *Encoder) AppendUint16(u uint16) {
	p.addElementSeparator()
	p.buf.AppendUint(uint64(u))
}

func (p *Encoder) AppendUint8(u uint8) {
	p.addElementSeparator()
	p.buf.AppendUint(uint64(u))
}

func (p *Encoder) AppendUintptr(u uintptr) {
	p.addElementSeparator()
	p.buf.AppendUint(uint64(u))
}

func (p *Encoder) AddArray(key string, marshaler zapcore.ArrayMarshaler) error {
	p.addKey(key)
	return marshaler.MarshalLogArray(p)
}

func (p *Encoder) AddObject(key string, marshaler zapcore.ObjectMarshaler) error {
	p.addKey(key)
	return marshaler.MarshalLogObject(p)
}

func (p *Encoder) AddBinary(key string, value []byte) {
	p.addKey(key)
	p.buf.AppendString(base64.StdEncoding.EncodeToString(value))
}

func (p *Encoder) AddByteString(key string, value []byte) {
	p.addKey(key)
	p.buf.AppendString(string(value))
}

func (p *Encoder) AddBool(key string, value bool) {
	p.addKey(key)
	p.buf.AppendBool(value)
}

func (p *Encoder) AddComplex128(key string, value complex128) {
	p.addKey(key)
	p.AppendComplex128(value)
}

func (p *Encoder) AddComplex64(key string, value complex64) {
	p.addKey(key)
	p.AddComplex128(key, complex128(value))
}

func (p *Encoder) AddDuration(key string, value time.Duration) {
	p.addKey(key)
	p.EncodeDuration(value, p)
}

func (p *Encoder) AddFloat64(key string, value float64) {
	p.addKey(key)
	p.buf.AppendFloat(value, 64)
}

func (p *Encoder) AddFloat32(key string, value float32) {
	p.addKey(key)
	p.buf.AppendFloat(float64(value), 64)
}

func (p *Encoder) AddInt(key string, value int) {
	p.AddInt64(key, int64(value))
}

func (p *Encoder) AddInt64(key string, value int64) {
	p.addKey(key)
	p.buf.AppendInt(value)
}

func (p *Encoder) AddInt32(key string, value int32) {
	p.AddInt64(key, int64(value))
}

func (p *Encoder) AddInt16(key string, value int16) {
	p.AddInt64(key, int64(value))
}

func (p *Encoder) AddInt8(key string, value int8) {
	p.AddInt64(key, int64(value))
}

func (p *Encoder) AddString(key, value string) {
	p.addKey(key)
	p.AppendString(value)
}

func (p *Encoder) AddTime(key string, value time.Time) {
	p.addKey(key)
	p.EncodeTime(value, p)
}

func (p *Encoder) AddUint(key string, value uint) {
	p.AddUint64(key, uint64(value))
}

func (p *Encoder) AddUint64(key string, value uint64) {
	p.addKey(key)
	p.AppendUint64(value)
}

func (p *Encoder) AddUint32(key string, value uint32) {
	p.AddUint64(key, uint64(value))
}

func (p *Encoder) AddUint16(key string, value uint16) {
	p.AddUint64(key, uint64(value))
}

func (p *Encoder) AddUint8(key string, value uint8) {
	p.AddUint64(key, uint64(value))
}

func (p *Encoder) AddUintptr(key string, value uintptr) {
	p.AddUint64(key, uint64(value))
}

func (p *Encoder) AddReflected(key string, value interface{}) error {
	p.addKey(key)
	v, ok := value.(string)
	if !ok {
		v = fmt.Sprintf("%v", value)
	}
	p.buf.AppendString(v)
	return nil
}

func (p *Encoder) OpenNamespace(key string) {}

func (p *Encoder) clone() *Encoder {
	bp := _bufferPool.Get()
	return &Encoder{
		EncoderConfig: p.EncoderConfig,
		buf:           bp,
	}
}

func (p *Encoder) Clone() zapcore.Encoder {
	return p.clone()
}

func (p *Encoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
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
