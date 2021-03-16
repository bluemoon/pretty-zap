package prettyzap

import (
	"sync"
)

var pool *sync.Pool

type EncoderPool struct{}

func NewEncoderPool() *EncoderPool {
	pool = &sync.Pool{
		New: func() interface{} {
			return &Encoder{}
		},
	}

	return &EncoderPool{}
}

func (e *EncoderPool) Get() *Encoder {
	return pool.Get().(*Encoder)
}

func (e *EncoderPool) Put(enc *Encoder) {
	enc.config = nil
	enc.buf = nil
	enc.termLastAppended = false
	pool.Put(enc)
}
