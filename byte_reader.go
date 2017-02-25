package kproto

import (
	"io"
	"unsafe"
)

type ByteReader struct {
	buf []byte
	l   int
	off int
}

func (b *ByteReader) Reset(buf []byte) {
	b.buf = buf
	b.l = len(buf)
	b.off = 0
}

func (b *ByteReader) Bytes() []byte {
	return b.buf
}

func (b *ByteReader) Length() int {
	return b.l
}

func (b *ByteReader) Offset() int {
	return b.off
}

func (b *ByteReader) Read(p []byte) error {
	l := len(p)
	if b.l-b.off < l {
		return io.EOF
	}

	copy(p, b.buf[b.off:])
	b.off += l
	return nil
}

func (b *ByteReader) ReadLength() (v int, err error) {
	if !b.Check(2) {
		err = io.EOF
		return
	}
	v = int(order.Uint16(b.buf[b.off:]))
	b.off += 2
	return
}

func (b *ByteReader) ReadByte() (v byte, err error) {
	if !b.Check(1) {
		err = io.EOF
		return
	}
	v = b.buf[b.off]
	b.off += 1
	return
}

func (b *ByteReader) ReadByteArray() (v []byte, err error) {
	var l int
	l, err = b.ReadLength()
	if err != nil {
		return
	}
	if !b.Check(l) {
		err = io.EOF
		return
	}
	v = make([]byte, l)
	for i := 0; i < l; i++ {
		v[i] = b.buf[b.off+i]
	}
	b.off += l
	return
}

func (b *ByteReader) ReadBool() (v bool, err error) {
	if !b.Check(1) {
		err = io.EOF
		return
	}
	v = b.buf[b.off] != 0
	b.off += 1
	return
}

func (b *ByteReader) ReadBoolArray() (v []bool, err error) {
	var l int
	l, err = b.ReadLength()
	if err != nil {
		return
	}
	if !b.Check(l) {
		err = io.EOF
		return
	}
	v = make([]bool, l)
	for i := 0; i < l; i++ {
		v[i] = b.buf[b.off+i] != 0
	}
	b.off += l
	return
}

func (b *ByteReader) ReadInt8() (v int8, err error) {
	if !b.Check(1) {
		err = io.EOF
		return
	}
	v = int8(b.buf[b.off])
	b.off += 1
	return
}

func (b *ByteReader) ReadInt8Array() (v []int8, err error) {
	var l int
	l, err = b.ReadLength()
	if err != nil {
		return
	}
	if !b.Check(l) {
		err = io.EOF
		return
	}
	v = make([]int8, l)
	for i := 0; i < l; i++ {
		v[i] = int8(b.buf[b.off+i])
	}
	b.off += l
	return
}

func (b *ByteReader) ReadUint8() (v uint8, err error) {
	if !b.Check(1) {
		err = io.EOF
		return
	}
	v = b.buf[b.off]
	b.off += 1
	return
}

func (b *ByteReader) ReadUint8Array() (v []uint8, err error) {
	var l int
	l, err = b.ReadLength()
	if err != nil {
		return
	}
	if !b.Check(l) {
		err = io.EOF
		return
	}
	v = make([]uint8, l)
	for i := 0; i < l; i++ {
		v[i] = b.buf[b.off+i]
	}
	b.off += l
	return
}

func (b *ByteReader) ReadInt16() (v int16, err error) {
	if !b.Check(2) {
		err = io.EOF
		return
	}
	v = int16(order.Uint16(b.buf[b.off:]))
	b.off += 2
	return
}

func (b *ByteReader) ReadInt16Array() (v []int16, err error) {
	var l int
	l, err = b.ReadLength()
	if err != nil {
		return
	}
	if !b.Check(l * 2) {
		err = io.EOF
		return
	}
	v = make([]int16, l)
	for i := 0; i < l; i++ {
		v[i] = int16(order.Uint16(b.buf[b.off+i*2:]))
	}
	b.off += l * 2
	return
}

func (b *ByteReader) ReadUint16() (v uint16, err error) {
	if !b.Check(2) {
		err = io.EOF
		return
	}
	v = order.Uint16(b.buf[b.off:])
	b.off += 2
	return
}

func (b *ByteReader) ReadUint16Array() (v []uint16, err error) {
	var l int
	l, err = b.ReadLength()
	if err != nil {
		return
	}
	if !b.Check(l * 2) {
		err = io.EOF
		return
	}
	v = make([]uint16, l)
	for i := 0; i < l; i++ {
		v[i] = order.Uint16(b.buf[b.off+i*2:])
	}
	b.off += l * 2
	return
}

func (b *ByteReader) ReadInt32() (v int32, err error) {
	if !b.Check(4) {
		err = io.EOF
		return
	}
	v = int32(order.Uint32(b.buf[b.off:]))
	b.off += 4
	return
}

func (b *ByteReader) ReadInt32Array() (v []int32, err error) {
	var l int
	l, err = b.ReadLength()
	if err != nil {
		return
	}
	if !b.Check(l * 4) {
		err = io.EOF
		return
	}
	v = make([]int32, l)
	for i := 0; i < l; i++ {
		v[i] = int32(order.Uint32(b.buf[b.off+i*4:]))
	}
	b.off += l * 4
	return
}

func (b *ByteReader) ReadUint32() (v uint32, err error) {
	if !b.Check(4) {
		err = io.EOF
		return
	}
	v = order.Uint32(b.buf[b.off:])
	b.off += 4
	return
}

func (b *ByteReader) ReadUint32Array() (v []uint32, err error) {
	var l int
	l, err = b.ReadLength()
	if err != nil {
		return
	}
	if !b.Check(l * 4) {
		err = io.EOF
		return
	}
	v = make([]uint32, l)
	for i := 0; i < l; i++ {
		v[i] = order.Uint32(b.buf[b.off+i*4:])
	}
	b.off += l * 4
	return
}

func (b *ByteReader) ReadInt64() (v int64, err error) {
	if !b.Check(8) {
		err = io.EOF
		return
	}
	v = int64(order.Uint64(b.buf[b.off:]))
	b.off += 8
	return
}

func (b *ByteReader) ReadInt64Array() (v []int64, err error) {
	var l int
	l, err = b.ReadLength()
	if err != nil {
		return
	}
	if !b.Check(l * 8) {
		err = io.EOF
		return
	}
	v = make([]int64, l)
	for i := 0; i < l; i++ {
		v[i] = int64(order.Uint64(b.buf[b.off+i*8:]))
	}
	b.off += l * 8
	return
}

func (b *ByteReader) ReadUint64() (v uint64, err error) {
	if !b.Check(8) {
		err = io.EOF
		return
	}
	v = order.Uint64(b.buf[b.off:])
	b.off += 8
	return
}

func (b *ByteReader) ReadUint64Array() (v []uint64, err error) {
	var l int
	l, err = b.ReadLength()
	if err != nil {
		return
	}
	if !b.Check(l * 8) {
		err = io.EOF
		return
	}
	v = make([]uint64, l)
	for i := 0; i < l; i++ {
		v[i] = order.Uint64(b.buf[b.off+i*8:])
	}
	b.off += l * 8
	return
}

func (b *ByteReader) ReadFloat32() (v float32, err error) {
	var u32 uint32
	u32, err = b.ReadUint32()
	if err != nil {
		return
	}
	v = *(*float32)(unsafe.Pointer(&u32))
	return
}

func (b *ByteReader) ReadFloat32Array() (v []float32, err error) {
	var l int
	l, err = b.ReadLength()
	if err != nil {
		return
	}
	if !b.Check(l * 4) {
		err = io.EOF
		return
	}
	v = make([]float32, l)
	for i := 0; i < l; i++ {
		v[i], _ = b.ReadFloat32()
	}
	return
}

func (b *ByteReader) ReadFloat64() (v float64, err error) {
	var u64 uint64
	u64, err = b.ReadUint64()
	if err != nil {
		return
	}
	v = *(*float64)(unsafe.Pointer(&u64))
	return
}

func (b *ByteReader) ReadFloat64Array() (v []float64, err error) {
	var l int
	l, err = b.ReadLength()
	if err != nil {
		return
	}
	if !b.Check(l * 8) {
		err = io.EOF
		return
	}
	v = make([]float64, l)
	for i := 0; i < l; i++ {
		v[i], _ = b.ReadFloat64()
	}
	return
}

func (b *ByteReader) ReadString() (v string, err error) {
	var l int
	l, err = b.ReadLength()
	if err != nil {
		return
	}
	if !b.Check(l) {
		err = io.EOF
		return
	}
	buf := make([]byte, l)
	copy(buf, b.buf)
	v = string(buf)
	return
}

func (b *ByteReader) ReadStringArray() (v []string, err error) {
	var l int
	l, err = b.ReadLength()
	if err != nil {
		return
	}
	v = make([]string, l)
	for i := 0; i < l; i++ {
		v[i], err = b.ReadString()
		if err != nil {
			return
		}
	}
	return
}

func (b *ByteReader) Check(n int) bool {
	if b.l-b.off < n {
		return false
	}

	return true
}

func NewByteReader(buf []byte) *ByteReader {
	return &ByteReader{buf, len(buf), 0}
}
