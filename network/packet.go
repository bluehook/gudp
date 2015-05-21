// Copyright 2015 The GUDP Authors. All rights reserved.
// HTTPS clone URL: https://github.com/bluehook/gudp.git

package network

import (
	"bytes"
	"encoding/binary"
)

//#底层数据包
// 主要作用是序列化上层数据
// 需要传入一个*bytes.Buffer作为缓存区
type Packet struct {
	buf   *bytes.Buffer    //缓存区
	order binary.ByteOrder //字节序
}

func NewPacket(order binary.ByteOrder) *Packet {
	packet := &Packet{order: order}
	return packet
}

func (self *Packet) SetBuf(buf *bytes.Buffer) {
	self.buf = buf
}

func (self *Packet) GetBuf() *bytes.Buffer {
	return self.buf
}

//##序列化
func (self *Packet) WriteByte(b byte) error {
	return binary.Write(self.buf, self.order, b)
}

func (self *Packet) WriteString(b string) error {
	err := binary.Write(self.buf, self.order, uint16(len(b)))
	if err != nil {
		return err
	}
	return binary.Write(self.buf, self.order, []byte(b))
}

func (self *Packet) WriteAny(b interface{}) error {
	return binary.Write(self.buf, self.order, b)
}

func (self *Packet) WriteInt8(b int8) error {
	return binary.Write(self.buf, self.order, b)
}

func (self *Packet) WriteUInt8(b uint8) error {
	return binary.Write(self.buf, self.order, b)
}

func (self *Packet) WriteInt16(b int16) error {
	return binary.Write(self.buf, self.order, b)
}

func (self *Packet) WriteUInt16(b uint16) error {
	return binary.Write(self.buf, self.order, b)
}

func (self *Packet) WriteInt32(b int32) error {
	return binary.Write(self.buf, self.order, b)
}

func (self *Packet) WriteUInt32(b uint32) error {
	return binary.Write(self.buf, self.order, b)
}

func (self *Packet) WriteInt64(b int64) error {
	return binary.Write(self.buf, self.order, b)
}

func (self *Packet) WriteUInt64(b uint64) error {
	return binary.Write(self.buf, self.order, b)
}

func (self *Packet) WriteFloat32(b float32) error {
	return binary.Write(self.buf, self.order, b)
}

func (self *Packet) WriteFloat64(b float64) error {
	return binary.Write(self.buf, self.order, b)
}

//##反序列化
func (self *Packet) ReadByte(b *byte) error {
	return binary.Read(self.buf, self.order, b)
}

func (self *Packet) ReadString(b *string) error {
	var size uint16
	err := binary.Read(self.buf, self.order, &size)
	if err != nil {
		return err
	}
	s := make([]byte, size)
	err = binary.Read(self.buf, self.order, &s)
	if err == nil {
		*b = string(s)
	}
	return err
}

func (self *Packet) ReadAny(b interface{}) error {
	return binary.Read(self.buf, self.order, b)
}

func (self *Packet) ReadInt8(b *int8) error {
	return binary.Read(self.buf, self.order, b)
}

func (self *Packet) ReadUInt8(b *uint8) error {
	return binary.Read(self.buf, self.order, b)
}

func (self *Packet) ReadInt16(b *int16) error {
	return binary.Read(self.buf, self.order, b)
}

func (self *Packet) ReadUInt16(b *uint16) error {
	return binary.Read(self.buf, self.order, b)
}

func (self *Packet) ReadInt32(b *int32) error {
	return binary.Read(self.buf, self.order, b)
}

func (self *Packet) ReadUInt32(b *uint32) error {
	return binary.Read(self.buf, self.order, b)
}

func (self *Packet) ReadInt64(b *int64) error {
	return binary.Read(self.buf, self.order, b)
}

func (self *Packet) ReadUInt64(b *uint64) error {
	return binary.Read(self.buf, self.order, b)
}

func (self *Packet) ReadFloat32(b *float32) error {
	return binary.Read(self.buf, self.order, b)
}

func (self *Packet) ReadFloat64(b *float64) error {
	return binary.Read(self.buf, self.order, b)
}
