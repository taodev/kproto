package kproto

import (
	"io"
	"unsafe"
)

type ByteWriter struct {
	buf []byte
	l   int
	off int
}

func (b *ByteWriter) Reset(buf []byte) {
	b.buf = buf
	b.l = len(buf)
	b.off = 0
}

func (b *ByteWriter) Bytes() []byte {
	return b.buf
}

func (b *ByteWriter) Length() int {
	return b.l
}

func (b *ByteWriter) Offset() int {
	return b.off
}

func (b *ByteWriter) Write(p []byte) error {
	l := len(p)
	if b.l-b.off < l {
		return io.EOF
	}

	copy(b.buf[b.off:], p)
	b.off += l
	return nil
}

func (b *ByteWriter) WriteLength(l int) error {
	if !b.Check(2) {
		return io.EOF
	}
	order.PutUint16(b.buf[b.off:], uint16(l))
	b.off += 2
	return nil
}

func (b *ByteWriter) WriteByte(v byte) error {
	if !b.Check(1) {
		return io.EOF
	}
	b.buf[b.off] = v
	b.off += 1
	return nil
}

func (b *ByteWriter) WriteByteArray(v []byte) error {
	l := len(v)
	if !b.Check(l + 2) {
		return io.EOF
	}
	b.WriteLength(l)
	copy(b.buf[b.off:], v)
	b.off += l
	return nil
}

func (b *ByteWriter) WriteBool(v bool) error {
	if !b.Check(1) {
		return io.EOF
	}
	if v {
		b.buf[b.off] = 1
	} else {
		b.buf[b.off] = 0
	}
	b.off += 1
	return nil
}

func (b *ByteWriter) WriteBoolArray(v []bool) error {
	l := len(v)
	if !b.Check(l + 2) {
		return io.EOF
	}
	b.WriteLength(l)
	for i := 0; i < l; i++ {
		if v[i] {
			b.buf[b.off+i] = 1
		} else {
			b.buf[b.off+i] = 0
		}
	}
	b.off += l
	return nil
}

func (b *ByteWriter) WriteInt8(v int8) error {
	if !b.Check(1) {
		return io.EOF
	}
	b.buf[b.off] = byte(v)
	b.off += 1
	return nil
}

func (b *ByteWriter) WriteInt8Array(v []int8) error {
	l := len(v)
	if !b.Check(l + 2) {
		return io.EOF
	}
	b.WriteLength(l)
	for i := 0; i < l; i++ {
		b.buf[b.off+i] = byte(v[i])
	}
	b.off += l
	return nil
}

func (b *ByteWriter) WriteUint8(v uint8) error {
	if !b.Check(1) {
		return io.EOF
	}
	b.buf[b.off] = v
	b.off += 1
	return nil
}

func (b *ByteWriter) WriteUint8Array(v []uint8) error {
	l := len(v)
	if !b.Check(l + 2) {
		return io.EOF
	}
	b.WriteLength(l)
	for i := 0; i < l; i++ {
		b.buf[b.off+i] = v[i]
	}
	b.off += l
	return nil
}

func (b *ByteWriter) WriteInt16(v int16) error {
	if !b.Check(2) {
		return io.EOF
	}
	order.PutUint16(b.buf[b.off:], uint16(v))
	b.off += 2
	return nil
}

func (b *ByteWriter) WriteInt16Array(v []int16) error {
	l := len(v)
	if !b.Check(l*2 + 2) {
		return io.EOF
	}
	b.WriteLength(l)
	for i := 0; i < l; i++ {
		order.PutUint16(b.buf[b.off+i*2:], uint16(v[i]))
	}
	b.off += l * 2
	return nil
}

func (b *ByteWriter) WriteUint16(v uint16) error {
	if !b.Check(2) {
		return io.EOF
	}
	order.PutUint16(b.buf[b.off:], v)
	b.off += 2
	return nil
}

func (b *ByteWriter) WriteUint16Array(v []uint16) error {
	l := len(v)
	if !b.Check(l*2 + 2) {
		return io.EOF
	}
	b.WriteLength(l)
	for i := 0; i < l; i++ {
		order.PutUint16(b.buf[b.off+i*2:], v[i])
	}
	b.off += l * 2
	return nil
}

func (b *ByteWriter) WriteInt32(v int32) error {
	if !b.Check(4) {
		return io.EOF
	}
	order.PutUint32(b.buf[b.off:], uint32(v))
	b.off += 4
	return nil
}

func (b *ByteWriter) WriteInt32Array(v []int32) error {
	l := len(v)
	if !b.Check(l*4 + 2) {
		return io.EOF
	}
	b.WriteLength(l)
	for i := 0; i < l; i++ {
		order.PutUint32(b.buf[b.off+i*4:], uint32(v[i]))
	}
	b.off += l * 4
	return nil
}

func (b *ByteWriter) WriteUint32(v uint32) error {
	if !b.Check(4) {
		return io.EOF
	}
	order.PutUint32(b.buf[b.off:], v)
	b.off += 4
	return nil
}

func (b *ByteWriter) WriteUint32Array(v []uint32) error {
	l := len(v)
	if !b.Check(l*4 + 2) {
		return io.EOF
	}
	b.WriteLength(l)
	for i := 0; i < l; i++ {
		order.PutUint32(b.buf[b.off+i*4:], v[i])
	}
	b.off += l * 4
	return nil
}

func (b *ByteWriter) WriteInt64(v int64) error {
	if !b.Check(8) {
		return io.EOF
	}
	order.PutUint64(b.buf[b.off:], uint64(v))
	b.off += 8
	return nil
}

func (b *ByteWriter) WriteInt64Array(v []int64) error {
	l := len(v)
	if !b.Check(l*8 + 2) {
		return io.EOF
	}
	b.WriteLength(l)
	for i := 0; i < l; i++ {
		order.PutUint64(b.buf[b.off+i*8:], uint64(v[i]))
	}
	b.off += l * 8
	return nil
}

func (b *ByteWriter) WriteUint64(v uint64) error {
	if !b.Check(8) {
		return io.EOF
	}
	order.PutUint64(b.buf[b.off:], v)
	b.off += 8
	return nil
}

func (b *ByteWriter) WriteUint64Array(v []uint64) error {
	l := len(v)
	if !b.Check(l*8 + 2) {
		return io.EOF
	}
	b.WriteLength(l)
	for i := 0; i < l; i++ {
		order.PutUint64(b.buf[b.off+i*8:], v[i])
	}
	b.off += l * 8
	return nil
}

func (b *ByteWriter) WriteFloat32(v float32) error {
	return b.WriteUint32(*(*uint32)(unsafe.Pointer(&v)))
}

func (b *ByteWriter) WriteFloat32Array(v []float32) error {
	l := len(v)
	if !b.Check(l*4 + 2) {
		return io.EOF
	}
	b.WriteLength(l)
	for _, x := range v {
		b.WriteUint32(*(*uint32)(unsafe.Pointer(&x)))
	}
	return nil
}

func (b *ByteWriter) WriteFloat64(v float64) error {
	return b.WriteUint64(*(*uint64)(unsafe.Pointer(&v)))
}

func (b *ByteWriter) WriteFloat64Array(v []float64) error {
	l := len(v)
	if !b.Check(l*8 + 2) {
		return io.EOF
	}
	b.WriteLength(l)
	for _, x := range v {
		b.WriteUint64(*(*uint64)(unsafe.Pointer(&x)))
	}
	return nil
}

func (b *ByteWriter) WriteString(v string) error {
	l := len(v)
	if !b.Check(l + 2) {
		return io.EOF
	}
	b.WriteLength(l)
	b.Write([]byte(v))
	return nil
}

func (b *ByteWriter) WriteStringArray(v []string) error {
	l := len(v)
	var err error
	if err = b.WriteLength(l); err != nil {
		return err
	}
	for i := 0; i < l; i++ {
		if err = b.WriteString(v[i]); err != nil {
			return err
		}
	}
	return nil
}

func (b *ByteWriter) Check(n int) bool {
	if b.l-b.off < n {
		return false
	}

	return true
}

func NewByteWriter(buf []byte) *ByteWriter {
	return &ByteWriter{buf, len(buf), 0}
}
