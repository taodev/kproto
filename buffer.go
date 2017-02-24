package kproto

import (
	"bytes"
	"errors"
	"io"
	"unsafe"
)

var (
	errWriteFailed = errors.New("Buffer write failed")
)

type Buffer struct {
	bytes.Buffer
	cache [8]byte
}

type MyBuffer struct {
	buf []byte
	off int
}

func NewBufferSize(s int) *MyBuffer {
	b := &MyBuffer{}
	b.buf = make([]byte, s)
	b.off = 0
	return b
}

func NewBuffer(buf []byte) *MyBuffer {
	b := &MyBuffer{}
	b.buf = buf
	b.off = 0
	return b
}

func (b *MyBuffer) Write(p []byte) error {
	l := len(p)
	if len(b.buf)-b.off < l {
		return io.EOF
	}

	copy(b.buf[b.off:], p)
	b.off += l
	return nil
}

func (b *MyBuffer) Read(p []byte) error {
	l := len(p)
	if len(b.buf)-b.off < l {
		return io.EOF
	}

	copy(p, b.buf[b.off:])
	b.off += l
	return nil
}

func (b *MyBuffer) WriteByte(v byte) error {
	if len(b.buf)-b.off < 1 {
		return io.EOF
	}

	b.buf[b.off] = v
	b.off += 1
	return nil
}

func (b *MyBuffer) ReadByte(v *byte) error {
	if len(b.buf)-b.off < 1 {
		return io.EOF
	}

	*v = b.buf[b.off]
	b.off += 1
	return nil
}

func (b *MyBuffer) WriteByteArray(v []byte) error {
	return b.Write(v)
}

func (b *MyBuffer) ReadByteArray(v []byte) error {
	return b.Read(v)
}

func (b *MyBuffer) WriteBool(v bool) error {
	if v {
		return b.WriteByte(1)
	}

	return b.WriteByte(0)
}

func (b *MyBuffer) ReadBool(v *bool) error {
	if len(b.buf)-b.off < 1 {
		return io.EOF
	}

	*v = b.buf[b.off] != 0
	b.off += 1
	return nil
}

func (b *MyBuffer) WriteBoolArray(v []bool) error {
	l := len(v)
	if len(b.buf)-b.off < l {
		return io.EOF
	}

	for i, x := range v {
		if x {
			b.buf[b.off+i] = 1
		} else {
			b.buf[b.off+i] = 0
		}
	}
	b.off += l
	return nil
}

func (b *MyBuffer) ReadBoolArray(v []bool) error {
	l := len(v)
	if len(b.buf)-b.off < l {
		return io.EOF
	}

	for i := 0; i < l; i++ {
		v[i] = b.buf[b.off+i] != 0
	}
	b.off += l
	return nil
}

func (b *MyBuffer) WriteInt8(v int8) error {
	if len(b.buf)-b.off < 1 {
		return io.EOF
	}
	b.buf[b.off] = byte(v)
	b.off += 1
	return nil
}

func (b *MyBuffer) ReadInt8(v *int8) error {
	if len(b.buf)-b.off < 1 {
		return io.EOF
	}
	*v = int8(b.buf[b.off])
	b.off += 1
	return nil
}

func (b *MyBuffer) WriteInt8Array(v []int8) error {
	l := len(v)
	if len(b.buf)-b.off < l {
		return io.EOF
	}
	for i, x := range v {
		b.buf[b.off+i] = byte(x)
	}
	b.off += l
	return nil
}

func (b *MyBuffer) ReadInt8Array(v []int8) error {
	l := len(v)
	if len(b.buf)-b.off < l {
		return io.EOF
	}
	for i := 0; i < l; i++ {
		v[i] = int8(b.buf[b.off+i])
	}
	b.off += l
	return nil
}

func (b *MyBuffer) WriteUint8(v uint8) error {
	return b.WriteByte(v)
}

func (b *MyBuffer) ReadUint8(v *uint8) error {
	return b.ReadByte(v)
}

func (b *MyBuffer) WriteUint8Array(v []uint8) error {
	return b.Write(v)
}

func (b *MyBuffer) ReadUint8Array(v []uint8) error {
	return b.Read(v)
}

func (b *MyBuffer) WriteInt16(v int16) error {
	if len(b.buf)-b.off < 2 {
		return io.EOF
	}
	b.buf[b.off] = byte(v)
	b.buf[b.off+1] = byte(v >> 8)
	b.off += 2
	return nil
}

func (b *MyBuffer) ReadInt16(v *int16) error {
	if len(b.buf)-b.off < 2 {
		return io.EOF
	}
	*v = int16(b.buf[b.off]) | int16(b.buf[b.off+1])<<8
	b.off += 2
	return nil
}

func (b *MyBuffer) WriteInt16Array(v []int16) error {
	l := len(v)
	if len(b.buf)-b.off < l*2 {
		return io.EOF
	}
	for i, x := range v {
		b.buf[b.off+i*2] = byte(x)
		b.buf[b.off+i*2+1] = byte(x >> 8)
	}
	b.off += l * 2
	return nil
}

func (b *MyBuffer) ReadInt16Array(v []int16) error {
	l := len(v)
	if len(b.buf)-b.off < l*2 {
		return io.EOF
	}
	for i := range v {
		v[i] = int16(b.buf[b.off+i*2]) | int16(b.buf[b.off+ +i*2+1])<<8
	}
	b.off += l * 2
	return nil
}

func (b *MyBuffer) WriteUint16(v uint16) error {
	if len(b.buf)-b.off < 2 {
		return io.EOF
	}
	b.buf[b.off] = byte(v)
	b.buf[b.off+1] = byte(v >> 8)
	b.off += 2
	return nil
}

func (b *MyBuffer) ReadUint16(v *uint16) error {
	if len(b.buf)-b.off < 2 {
		return io.EOF
	}
	*v = uint16(b.buf[b.off]) | uint16(b.buf[b.off+1])<<8
	b.off += 2
	return nil
}

func (b *MyBuffer) WriteUint16Array(v []uint16) error {
	l := len(v)
	if len(b.buf)-b.off < l*2 {
		return io.EOF
	}
	for i, x := range v {
		b.buf[b.off+i*2] = byte(x)
		b.buf[b.off+i*2+1] = byte(x >> 8)
	}
	b.off += l * 2
	return nil
}

func (b *MyBuffer) ReadUint16Array(v []uint16) error {
	l := len(v)
	if len(b.buf)-b.off < l*2 {
		return io.EOF
	}
	for i := range v {
		v[i] = uint16(b.buf[b.off+i*2]) | uint16(b.buf[b.off+ +i*2+1])<<8
	}
	b.off += l * 2
	return nil
}

func (b *MyBuffer) WriteInt32(v int32) error {
	if len(b.buf)-b.off < 4 {
		return io.EOF
	}
	b.buf[b.off] = byte(v)
	b.buf[b.off+1] = byte(v >> 8)
	b.buf[b.off+2] = byte(v >> 16)
	b.buf[b.off+3] = byte(v >> 24)
	b.off += 4
	return nil
}

func (b *MyBuffer) ReadInt32(v *int32) error {
	if len(b.buf)-b.off < 4 {
		return io.EOF
	}
	*v = int32(b.buf[b.off]) | int32(b.buf[b.off+1])<<8 |
		int32(b.buf[b.off+2])<<16 | int32(b.buf[b.off+3])<<24
	b.off += 4
	return nil
}

func (b *MyBuffer) WriteInt32Array(v []int32) error {
	l := len(v) * 4
	if len(b.buf)-b.off < l {
		return io.EOF
	}
	for i, x := range v {
		b.buf[b.off+i*4] = byte(x)
		b.buf[b.off+i*4+1] = byte(x >> 8)
		b.buf[b.off+i*4+2] = byte(x >> 16)
		b.buf[b.off+i*4+3] = byte(x >> 24)
	}
	b.off += l
	return nil
}

func (b *MyBuffer) ReadInt32Array(v []int32) error {
	l := len(v) * 4
	if len(b.buf)-b.off < l {
		return io.EOF
	}
	for i := range v {
		v[i] = int32(b.buf[b.off+i*4]) | int32(b.buf[b.off+i*4+1])<<8 |
			int32(b.buf[b.off+i*4+2])<<16 | int32(b.buf[b.off+i*4+3])<<24
	}
	b.off += l
	return nil
}

func (b *MyBuffer) WriteUint32(v uint32) error {
	if len(b.buf)-b.off < 4 {
		return io.EOF
	}
	b.buf[b.off] = byte(v)
	b.buf[b.off+1] = byte(v >> 8)
	b.buf[b.off+2] = byte(v >> 16)
	b.buf[b.off+3] = byte(v >> 24)
	b.off += 4
	return nil
}

func (b *MyBuffer) ReadUint32(v *uint32) error {
	if len(b.buf)-b.off < 4 {
		return io.EOF
	}
	*v = uint32(b.buf[b.off]) | uint32(b.buf[b.off+1])<<8 |
		uint32(b.buf[b.off+2])<<16 | uint32(b.buf[b.off+3])<<24
	b.off += 4
	return nil
}

func (b *MyBuffer) WriteUint32Array(v []uint32) error {
	l := len(v) * 4
	if len(b.buf)-b.off < l {
		return io.EOF
	}
	for i, x := range v {
		b.buf[b.off+i*4] = byte(x)
		b.buf[b.off+i*4+1] = byte(x >> 8)
		b.buf[b.off+i*4+2] = byte(x >> 16)
		b.buf[b.off+i*4+3] = byte(x >> 24)
	}
	b.off += l
	return nil
}

func (b *MyBuffer) ReadUint32Array(v []uint32) error {
	l := len(v) * 4
	if len(b.buf)-b.off < l {
		return io.EOF
	}
	for i := range v {
		v[i] = uint32(b.buf[b.off+i*4]) | uint32(b.buf[b.off+i*4+1])<<8 |
			uint32(b.buf[b.off+i*4+2])<<16 | uint32(b.buf[b.off+i*4+3])<<24
	}
	b.off += l
	return nil
}

func (b *MyBuffer) WriteInt64(v int64) error {
	if len(b.buf)-b.off < 8 {
		return io.EOF
	}
	b.buf[b.off] = byte(v)
	b.buf[b.off+1] = byte(v >> 8)
	b.buf[b.off+2] = byte(v >> 16)
	b.buf[b.off+3] = byte(v >> 24)
	b.buf[b.off+4] = byte(v >> 32)
	b.buf[b.off+5] = byte(v >> 40)
	b.buf[b.off+6] = byte(v >> 48)
	b.buf[b.off+7] = byte(v >> 56)
	b.off += 8
	return nil
}

func (b *MyBuffer) ReadInt64(v *int64) error {
	if len(b.buf)-b.off < 8 {
		return io.EOF
	}
	*v = int64(b.buf[b.off]) | int64(b.buf[b.off+1])<<8 |
		int64(b.buf[b.off+2])<<16 | int64(b.buf[b.off+3])<<24 |
		int64(b.buf[b.off+4])<<32 | int64(b.buf[b.off+5])<<40 |
		int64(b.buf[b.off+6])<<48 | int64(b.buf[b.off+7])<<56
	b.off += 8
	return nil
}

func (b *MyBuffer) WriteInt64rray(v []int64) error {
	l := len(v) * 8
	if len(b.buf)-b.off < l {
		return io.EOF
	}
	for i, x := range v {
		b.buf[b.off+i*8] = byte(x)
		b.buf[b.off+i*8+1] = byte(x >> 8)
		b.buf[b.off+i*8+2] = byte(x >> 16)
		b.buf[b.off+i*8+3] = byte(x >> 24)
		b.buf[b.off+i*8+4] = byte(x >> 32)
		b.buf[b.off+i*8+5] = byte(x >> 40)
		b.buf[b.off+i*8+6] = byte(x >> 48)
		b.buf[b.off+i*8+7] = byte(x >> 56)
	}
	b.off += l
	return nil
}

func (b *MyBuffer) ReadInt64Array(v []int64) error {
	l := len(v) * 8
	if len(b.buf)-b.off < l {
		return io.EOF
	}
	for i := range v {
		v[i] = int64(b.buf[b.off+i*8]) | int64(b.buf[b.off+i*8+1])<<8 |
			int64(b.buf[b.off+i*8+2])<<16 | int64(b.buf[b.off+i*8+3])<<24 |
			int64(b.buf[b.off+i*8+4])<<32 | int64(b.buf[b.off+i*8+5])<<40 |
			int64(b.buf[b.off+i*8+6])<<48 | int64(b.buf[b.off+i*8+7])<<56
	}
	b.off += l
	return nil
}

func (b *MyBuffer) WriteUint64(v uint64) error {
	if len(b.buf)-b.off < 8 {
		return io.EOF
	}
	b.buf[b.off] = byte(v)
	b.buf[b.off+1] = byte(v >> 8)
	b.buf[b.off+2] = byte(v >> 16)
	b.buf[b.off+3] = byte(v >> 24)
	b.buf[b.off+4] = byte(v >> 32)
	b.buf[b.off+5] = byte(v >> 40)
	b.buf[b.off+6] = byte(v >> 48)
	b.buf[b.off+7] = byte(v >> 56)
	b.off += 8
	return nil
}

func (b *MyBuffer) ReadUint64(v *uint64) error {
	if len(b.buf)-b.off < 8 {
		return io.EOF
	}
	*v = uint64(b.buf[b.off]) | uint64(b.buf[b.off+1])<<8 |
		uint64(b.buf[b.off+2])<<16 | uint64(b.buf[b.off+3])<<24 |
		uint64(b.buf[b.off+4])<<32 | uint64(b.buf[b.off+5])<<40 |
		uint64(b.buf[b.off+6])<<48 | uint64(b.buf[b.off+7])<<56
	b.off += 8
	return nil
}

func (b *MyBuffer) WriteUint64rray(v []uint64) error {
	l := len(v) * 8
	if len(b.buf)-b.off < l {
		return io.EOF
	}
	for i, x := range v {
		b.buf[b.off+i*8] = byte(x)
		b.buf[b.off+i*8+1] = byte(x >> 8)
		b.buf[b.off+i*8+2] = byte(x >> 16)
		b.buf[b.off+i*8+3] = byte(x >> 24)
		b.buf[b.off+i*8+4] = byte(x >> 32)
		b.buf[b.off+i*8+5] = byte(x >> 40)
		b.buf[b.off+i*8+6] = byte(x >> 48)
		b.buf[b.off+i*8+7] = byte(x >> 56)
	}
	b.off += l
	return nil
}

func (b *MyBuffer) ReadUint64Array(v []uint64) error {
	l := len(v) * 8
	if len(b.buf)-b.off < l {
		return io.EOF
	}
	for i := range v {
		v[i] = uint64(b.buf[b.off+i*8]) | uint64(b.buf[b.off+i*8+1])<<8 |
			uint64(b.buf[b.off+i*8+2])<<16 | uint64(b.buf[b.off+i*8+3])<<24 |
			uint64(b.buf[b.off+i*8+4])<<32 | uint64(b.buf[b.off+i*8+5])<<40 |
			uint64(b.buf[b.off+i*8+6])<<48 | uint64(b.buf[b.off+i*8+7])<<56
	}
	b.off += l
	return nil
}

func (b *MyBuffer) WriteFloat32(v float32) error {
	return b.WriteUint32(*(*uint32)(unsafe.Pointer(&v)))
}

func (b *MyBuffer) ReadFloat32(v *float32) error {
	var t uint32
	if err := b.ReadUint32(&t); err != nil {
		return err
	}
	*v = *(*float32)(unsafe.Pointer(&t))
	return nil
}

func (b *MyBuffer) WriteFloat32Array(v []float32) error {
	var err error
	for _, x := range v {
		if err = b.WriteFloat32(x); err != nil {
			return err
		}
	}
	return nil
}

func (b *MyBuffer) ReadFloat32Array(v []float32) error {
	var err error
	var t float32
	for i := range v {
		if err = b.ReadFloat32(&t); err != nil {
			return err
		}
		v[i] = t
	}
	return nil
}

func (b *MyBuffer) WriteFloat64(v float64) error {
	return b.WriteUint64(*(*uint64)(unsafe.Pointer(&v)))
}

func (b *MyBuffer) ReadFloat64(v *float64) error {
	var t uint64
	if err := b.ReadUint64(&t); err != nil {
		return err
	}
	*v = *(*float64)(unsafe.Pointer(&t))
	return nil
}

func (b *MyBuffer) WriteString(v string) (err error) {
	if err = b.WriteUint16(uint16(len(v))); err != nil {
		return
	}
	err = b.Write([]byte(v))
	return
}

func (b *MyBuffer) ReadString(v *string) (err error) {
	var l uint16
	if err = b.ReadUint16(&l); err != nil {
		return
	}
	bits := make([]byte, l)
	if err = b.Read(bits); err != nil {
		return
	}
	*v = string(bits)
	return
}
