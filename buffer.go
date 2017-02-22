package kproto

import (
	"bytes"
	"errors"
	"unsafe"
)

var (
	errWriteFailed = errors.New("Buffer write failed")
)

type Buffer struct {
	bytes.Buffer
	cache [8]byte
}

type Writer interface {
	Write(w *Buffer) error
}

func (b *Buffer) WriteBool(v bool) error {
	if v {
		return b.WriteByte(1)
	}

	return b.WriteByte(0)
}

func (b *Buffer) WriteArrayBool(v []bool) error {
	var err error
	if err = b.WriteUint16(uint16(len(v))); err != nil {
		return err
	}

	for _, x := range v {
		if err = b.WriteBool(x); err != nil {
			return err
		}
	}

	return nil
}

func (b *Buffer) WriteInt8(v int8) error {
	return b.WriteByte(byte(v))
}

func (b *Buffer) WriteArrayInt8(v []int8) error {
	var err error
	if err = b.WriteUint16(uint16(len(v))); err != nil {
		return err
	}

	for _, x := range v {
		if err = b.WriteInt8(x); err != nil {
			return err
		}
	}

	return nil
}

func (b *Buffer) WriteUint8(v uint8) error {
	return b.WriteByte(v)
}

func (b *Buffer) WriteArrayUint8(v []uint8) error {
	var err error
	if err = b.WriteUint16(uint16(len(v))); err != nil {
		return err
	}

	for _, x := range v {
		if err = b.WriteUint8(x); err != nil {
			return err
		}
	}

	return nil
}

func (b *Buffer) WriteInt16(v int16) error {
	b.cache[0] = byte(v)
	b.cache[1] = byte(v >> 8)

	n, err := b.Write(b.cache[:2])
	if n != 2 {
		return errWriteFailed
	}
	return err
}

func (b *Buffer) WriteArrayInt16(v []int16) error {
	var err error
	if err = b.WriteUint16(uint16(len(v))); err != nil {
		return err
	}

	for _, x := range v {
		if err = b.WriteInt16(x); err != nil {
			return err
		}
	}

	return nil
}

func (b *Buffer) WriteUint16(v uint16) error {
	b.cache[0] = byte(v)
	b.cache[1] = byte(v >> 8)

	n, err := b.Write(b.cache[:2])
	if n != 2 {
		return errWriteFailed
	}
	return err
}

func (b *Buffer) WriteArrayUint16(v []uint16) error {
	var err error
	if err = b.WriteUint16(uint16(len(v))); err != nil {
		return err
	}

	for _, x := range v {
		if err = b.WriteUint16(x); err != nil {
			return err
		}
	}

	return nil
}

func (b *Buffer) WriteInt32(v int32) error {
	b.cache[0] = byte(v)
	b.cache[1] = byte(v >> 8)
	b.cache[2] = byte(v >> 16)
	b.cache[3] = byte(v >> 24)

	n, err := b.Write(b.cache[:4])
	if n != 4 {
		return errWriteFailed
	}
	return err
}

func (b *Buffer) WriteArrayInt32(v []int32) error {
	var err error
	if err = b.WriteUint16(uint16(len(v))); err != nil {
		return err
	}

	for _, x := range v {
		if err = b.WriteInt32(x); err != nil {
			return err
		}
	}

	return nil
}

func (b *Buffer) WriteUint32(v uint32) error {
	b.cache[0] = byte(v)
	b.cache[1] = byte(v >> 8)
	b.cache[2] = byte(v >> 16)
	b.cache[3] = byte(v >> 24)

	n, err := b.Write(b.cache[:4])
	if n != 4 {
		return errWriteFailed
	}
	return err
}

func (b *Buffer) WriteArrayUint32(v []uint32) error {
	var err error
	if err = b.WriteUint16(uint16(len(v))); err != nil {
		return err
	}

	for _, x := range v {
		if err = b.WriteUint32(x); err != nil {
			return err
		}
	}

	return nil
}

func (b *Buffer) WriteInt64(v int64) error {
	b.cache[0] = byte(v)
	b.cache[1] = byte(v >> 8)
	b.cache[2] = byte(v >> 16)
	b.cache[3] = byte(v >> 24)
	b.cache[4] = byte(v >> 32)
	b.cache[5] = byte(v >> 40)
	b.cache[6] = byte(v >> 48)
	b.cache[7] = byte(v >> 56)

	n, err := b.Write(b.cache[:8])
	if n != 8 {
		return errWriteFailed
	}
	return err
}

func (b *Buffer) WriteArrayInt64(v []int64) error {
	var err error
	if err = b.WriteUint16(uint16(len(v))); err != nil {
		return err
	}

	for _, x := range v {
		if err = b.WriteInt64(x); err != nil {
			return err
		}
	}

	return nil
}

func (b *Buffer) WriteUint64(v uint64) error {
	b.cache[0] = byte(v)
	b.cache[1] = byte(v >> 8)
	b.cache[2] = byte(v >> 16)
	b.cache[3] = byte(v >> 24)
	b.cache[4] = byte(v >> 32)
	b.cache[5] = byte(v >> 40)
	b.cache[6] = byte(v >> 48)
	b.cache[7] = byte(v >> 56)

	n, err := b.Write(b.cache[:8])
	if n != 8 {
		return errWriteFailed
	}
	return err
}

func (b *Buffer) WriteArrayUint64(v []uint64) error {
	var err error
	if err = b.WriteUint16(uint16(len(v))); err != nil {
		return err
	}

	for _, x := range v {
		if err = b.WriteUint64(x); err != nil {
			return err
		}
	}

	return nil
}

func (b *Buffer) WriteFloat32(v float32) error {
	return b.WriteUint32(*(*uint32)(unsafe.Pointer(&v)))
}

func (b *Buffer) WriteArrayFloat32(v []float32) error {
	var err error
	if err = b.WriteUint16(uint16(len(v))); err != nil {
		return err
	}

	for _, x := range v {
		if err = b.WriteFloat32(x); err != nil {
			return err
		}
	}

	return nil
}

func (b *Buffer) WriteFloat64(v float64) error {
	return b.WriteUint64(*(*uint64)(unsafe.Pointer(&v)))
}

func (b *Buffer) WriteArrayFloat64(v []float64) error {
	var err error
	if err = b.WriteUint16(uint16(len(v))); err != nil {
		return err
	}

	for _, x := range v {
		if err = b.WriteFloat64(x); err != nil {
			return err
		}
	}

	return nil
}

func (b *Buffer) WriteString(v string) error {
	var err error
	l := len(v)
	if err = b.WriteUint16(uint16(l)); err != nil {
		return err
	}

	n, err := b.Write([]byte(v))
	if n != l {
		return errWriteFailed
	}
	return err
}

func (b *Buffer) WriteArrayString(v []string) error {
	var err error
	l := len(v)
	if err = b.WriteUint16(uint16(l)); err != nil {
		return err
	}

	for _, x := range v {
		if err = b.WriteString(x); err != nil {
			return err
		}
	}

	return nil
}

func (b *Buffer) WriteStruct(v Writer) error {
	return v.Write(b)
}
