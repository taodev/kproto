package kproto

import (
	"encoding/binary"
	"io"
	"unsafe"
)

var order binary.ByteOrder = binary.LittleEndian

func SetEndian(newOrder binary.ByteOrder) {
	order = newOrder
}

func GetEndian() binary.ByteOrder {
	return order
}

type ByteBuffer struct {
	buf []byte
	off int
}

func NewBufferSize(s int) *ByteBuffer {
	b := &ByteBuffer{}
	b.buf = make([]byte, s)
	b.off = 0
	return b
}

func NewBuffer(buf []byte) *ByteBuffer {
	b := &ByteBuffer{}
	b.buf = buf
	b.off = 0
	return b
}

func (b *ByteBuffer) Write(p []byte) error {
	l := len(p)
	if len(b.buf)-b.off < l {
		return io.EOF
	}

	copy(b.buf[b.off:], p)
	b.off += l
	return nil
}

func (b *ByteBuffer) Read(p []byte) error {
	l := len(p)
	if len(b.buf)-b.off < l {
		return io.EOF
	}

	copy(p, b.buf[b.off:])
	b.off += l
	return nil
}

func (b *ByteBuffer) WriteLength(length int) error {
	return b.WriteUint16(uint16(length))
}

func (b *ByteBuffer) ReadLength() (length int, err error) {
	var l uint16
	if err = b.ReadUint16(&l); err != nil {
		return
	}
	length = int(l)
	return
}

func (b *ByteBuffer) WriteByte(v byte) error {
	if len(b.buf)-b.off < 1 {
		return io.EOF
	}

	b.buf[b.off] = v
	b.off += 1
	return nil
}

func (b *ByteBuffer) ReadByte(v *byte) error {
	if len(b.buf)-b.off < 1 {
		return io.EOF
	}

	*v = b.buf[b.off]
	b.off += 1
	return nil
}

func (b *ByteBuffer) WriteByteArray(v []byte) error {
	if err := b.WriteLength(len(v)); err != nil {
		return err
	}

	return b.Write(v)
}

func (b *ByteBuffer) ReadByteArray(v *[]byte) error {
	l, err := b.ReadLength()
	if err != nil {
		return err
	}
	*v = make([]byte, l)
	return b.Read(*v)
}

func (b *ByteBuffer) WriteBool(v bool) error {
	if v {
		return b.WriteByte(1)
	}

	return b.WriteByte(0)
}

func (b *ByteBuffer) ReadBool(v *bool) error {
	if len(b.buf)-b.off < 1 {
		return io.EOF
	}

	*v = b.buf[b.off] != 0
	b.off += 1
	return nil
}

func (b *ByteBuffer) WriteBoolArray(v []bool) error {
	l := len(v)
	if len(b.buf)-b.off < l+2 {
		return io.EOF
	}

	b.WriteLength(l)

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

func (b *ByteBuffer) ReadBoolArray(v *[]bool) error {
	l, err := b.ReadLength()
	if err != nil {
		return err
	}

	if len(b.buf)-b.off < int(l) {
		err = io.EOF
		return io.EOF
	}

	for i := 0; i < int(l); i++ {
		(*v)[i] = b.buf[b.off+i] != 0
	}
	b.off += int(l)
	return nil
}

func (b *ByteBuffer) WriteInt8(v int8) error {
	if len(b.buf)-b.off < 1 {
		return io.EOF
	}
	b.buf[b.off] = byte(v)
	b.off += 1
	return nil
}

func (b *ByteBuffer) ReadInt8(v *int8) error {
	if len(b.buf)-b.off < 1 {
		return io.EOF
	}
	*v = int8(b.buf[b.off])
	b.off += 1
	return nil
}

func (b *ByteBuffer) WriteInt8Array(v []int8) error {
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

func (b *ByteBuffer) ReadInt8Array(v []int8) error {
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

func (b *ByteBuffer) WriteUint8(v uint8) error {
	return b.WriteByte(v)
}

func (b *ByteBuffer) ReadUint8(v *uint8) error {
	return b.ReadByte(v)
}

func (b *ByteBuffer) WriteUint8Array(v []uint8) error {
	return b.Write(v)
}

func (b *ByteBuffer) ReadUint8Array(v []uint8) error {
	return b.Read(v)
}

func (b *ByteBuffer) WriteInt16(v int16) error {
	if len(b.buf)-b.off < 2 {
		return io.EOF
	}
	order.PutUint16(b.buf[b.off:], uint16(v))
	b.off += 2
	return nil
}

func (b *ByteBuffer) ReadInt16(v *int16) error {
	if len(b.buf)-b.off < 2 {
		return io.EOF
	}
	*v = int16(order.Uint16(b.buf[b.off:]))
	b.off += 2
	return nil
}

func (b *ByteBuffer) WriteInt16Array(v []int16) error {
	if len(b.buf)-b.off < len(v)*2 {
		return io.EOF
	}
	for _, x := range v {
		order.PutUint16(b.buf[b.off:], uint16(x))
		b.off += 2
	}
	return nil
}

func (b *ByteBuffer) ReadInt16Array(v []int16) error {
	if len(b.buf)-b.off < len(v)*2 {
		return io.EOF
	}
	for i := range v {
		v[i] = int16(order.Uint16(b.buf[b.off:]))
		b.off += 2
	}
	return nil
}

func (b *ByteBuffer) WriteUint16(v uint16) error {
	if len(b.buf)-b.off < 2 {
		return io.EOF
	}
	order.PutUint16(b.buf[b.off:], v)
	b.off += 2
	return nil
}

func (b *ByteBuffer) ReadUint16(v *uint16) error {
	if len(b.buf)-b.off < 2 {
		return io.EOF
	}
	*v = order.Uint16(b.buf[b.off:])
	b.off += 2
	return nil
}

func (b *ByteBuffer) WriteUint16Array(v []uint16) error {
	if len(b.buf)-b.off < len(v)*2 {
		return io.EOF
	}
	for _, x := range v {
		order.PutUint16(b.buf[b.off:], x)
		b.off += 2
	}
	return nil
}

func (b *ByteBuffer) ReadUint16Array(v []uint16) error {
	if len(b.buf)-b.off < len(v)*2 {
		return io.EOF
	}
	for i := range v {
		v[i] = order.Uint16(b.buf[b.off:])
		b.off += 2
	}
	return nil
}

func (b *ByteBuffer) WriteInt32(v int32) error {
	if len(b.buf)-b.off < 4 {
		return io.EOF
	}
	order.PutUint32(b.buf[b.off:], uint32(v))
	b.off += 4
	return nil
}

func (b *ByteBuffer) ReadInt32(v *int32) error {
	if len(b.buf)-b.off < 4 {
		return io.EOF
	}
	*v = int32(order.Uint32(b.buf[b.off:]))
	b.off += 4
	return nil
}

func (b *ByteBuffer) WriteInt32Array(v []int32) error {
	if len(b.buf)-b.off < len(v)*4 {
		return io.EOF
	}
	for _, x := range v {
		order.PutUint32(b.buf[b.off:], uint32(x))
		b.off += 4
	}
	return nil
}

func (b *ByteBuffer) ReadInt32Array(v []int32) error {
	if len(b.buf)-b.off < len(v)*4 {
		return io.EOF
	}
	for i := range v {
		v[i] = int32(order.Uint32(b.buf[b.off:]))
		b.off += 4
	}
	return nil
}

func (b *ByteBuffer) WriteUint32(v uint32) error {
	if len(b.buf)-b.off < 4 {
		return io.EOF
	}
	order.PutUint32(b.buf[b.off:], v)
	b.off += 4
	return nil
}

func (b *ByteBuffer) ReadUint32(v *uint32) error {
	if len(b.buf)-b.off < 4 {
		return io.EOF
	}
	*v = order.Uint32(b.buf[b.off:])
	b.off += 4
	return nil
}

func (b *ByteBuffer) WriteUint32Array(v []uint32) error {
	if len(b.buf)-b.off < len(v)*4 {
		return io.EOF
	}
	for _, x := range v {
		order.PutUint32(b.buf[b.off:], x)
		b.off += 4
	}
	return nil
}

func (b *ByteBuffer) ReadUint32Array(v []uint32) error {
	if len(b.buf)-b.off < len(v)*4 {
		return io.EOF
	}
	for i := range v {
		v[i] = order.Uint32(b.buf[b.off:])
		b.off += 4
	}
	return nil
}

func (b *ByteBuffer) WriteInt64(v int64) error {
	if len(b.buf)-b.off < 8 {
		return io.EOF
	}
	order.PutUint64(b.buf[b.off:], uint64(v))
	b.off += 8
	return nil
}

func (b *ByteBuffer) ReadInt64(v *int64) error {
	if len(b.buf)-b.off < 8 {
		return io.EOF
	}
	*v = int64(order.Uint64(b.buf[b.off:]))
	b.off += 8
	return nil
}

func (b *ByteBuffer) WriteInt64Array(v []int64) error {
	if len(b.buf)-b.off < len(v)*8 {
		return io.EOF
	}
	for _, x := range v {
		order.PutUint64(b.buf[b.off:], uint64(x))
		b.off += 8
	}
	return nil
}

func (b *ByteBuffer) ReadInt64Array(v []int64) error {
	if len(b.buf)-b.off < len(v)*8 {
		return io.EOF
	}
	for i := range v {
		v[i] = int64(order.Uint64(b.buf[b.off:]))
		b.off += 8
	}
	return nil
}

func (b *ByteBuffer) WriteUint64(v uint64) error {
	if len(b.buf)-b.off < 8 {
		return io.EOF
	}
	order.PutUint64(b.buf[b.off:], v)
	b.off += 8
	return nil
}

func (b *ByteBuffer) ReadUint64(v *uint64) error {
	if len(b.buf)-b.off < 8 {
		return io.EOF
	}
	*v = order.Uint64(b.buf[b.off:])
	b.off += 8
	return nil
}

func (b *ByteBuffer) WriteUint64Array(v []uint64) error {
	if len(b.buf)-b.off < len(v)*8 {
		return io.EOF
	}
	for _, x := range v {
		order.PutUint64(b.buf[b.off:], x)
		b.off += 8
	}
	return nil
}

func (b *ByteBuffer) ReadUint64Array(v []uint64) error {
	if len(b.buf)-b.off < len(v)*8 {
		return io.EOF
	}
	for i := range v {
		v[i] = order.Uint64(b.buf[b.off:])
		b.off += 8
	}
	return nil
}

func (b *ByteBuffer) WriteFloat32(v float32) error {
	return b.WriteUint32(*(*uint32)(unsafe.Pointer(&v)))
}

func (b *ByteBuffer) ReadFloat32(v *float32) error {
	var t uint32
	if err := b.ReadUint32(&t); err != nil {
		return err
	}
	*v = *(*float32)(unsafe.Pointer(&t))
	return nil
}

func (b *ByteBuffer) WriteFloat32Array(v []float32) error {
	var err error
	for _, x := range v {
		if err = b.WriteFloat32(x); err != nil {
			return err
		}
	}
	return nil
}

func (b *ByteBuffer) ReadFloat32Array(v []float32) error {
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

func (b *ByteBuffer) WriteFloat64(v float64) error {
	return b.WriteUint64(*(*uint64)(unsafe.Pointer(&v)))
}

func (b *ByteBuffer) ReadFloat64(v *float64) error {
	var t uint64
	if err := b.ReadUint64(&t); err != nil {
		return err
	}
	*v = *(*float64)(unsafe.Pointer(&t))
	return nil
}

func (b *ByteBuffer) WriteFloat64Array(v []float64) error {
	var err error
	for _, x := range v {
		if err = b.WriteFloat64(x); err != nil {
			return err
		}
	}
	return nil
}

func (b *ByteBuffer) ReadFloat64Array(v []float64) error {
	var err error
	var t float64
	for i := range v {
		if err = b.ReadFloat64(&t); err != nil {
			return err
		}
		v[i] = t
	}
	return nil
}

func (b *ByteBuffer) WriteString(v string) (err error) {
	if err = b.WriteUint16(uint16(len(v))); err != nil {
		return
	}
	err = b.Write([]byte(v))
	return
}

func (b *ByteBuffer) ReadString(v *string) (err error) {
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

func (b *ByteBuffer) Bytes() []byte {
	return b.buf[:b.off]
}

func (b *ByteBuffer) Buffer() []byte {
	return b.buf
}
