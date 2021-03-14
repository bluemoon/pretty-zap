package main

import "go.uber.org/zap/buffer"

func isEmpty(str string) bool {
	return len(str) < 1
}

func c128(buf *buffer.Buffer, c complex128) {
	r, i := float64(real(c)), float64(imag(c))
	buf.AppendFloat(r, 64)
	buf.AppendByte('+')
	buf.AppendFloat(i, 64)
	buf.AppendByte('i')
}
