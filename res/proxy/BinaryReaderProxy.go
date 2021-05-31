package main

import (
	"bytes"
	"encoding/binary"
)

const (
	SIZE_OF_BOOLEAN int = 1
	SIZE_OF_INT8    int = 1
	SIZE_OF_INT16   int = 2
	SIZE_OF_INT32   int = 4
	SIZE_OF_INT64   int = 8
	SIZE_OF_UINT8   int = 1
	SIZE_OF_UINT16  int = 2
	SIZE_OF_UINT32  int = 4
	SIZE_OF_UINT64  int = 8
	SIZE_OF_FLOAT32 int = 4
	SIZE_OF_FLOAT64 int = 8
)

type IBinaryReaderProxy interface {
	//----------------------
	getBytesAvailable() bool
	//单个数据读取----------------------
	readBoolean() bool
	readUInt8() uint8
	readUInt16() uint16
	readUInt32() uint32
	readUInt64() uint64
	readInt8() int8
	readInt16() int16
	readInt32() int32
	readInt64() int64
	readFloat32() float32
	readFloat64() float64
	readString() string
	//数组数据读取----------------------
	readBooleanArray() []bool
	readUInt8Array() []uint8
	readUInt16Array() []uint16
	readUInt32Array() []uint32
	readUInt64Array() []uint64
	readInt8Array() []int8
	readInt16Array() []int16
	readInt32Array() []int32
	readInt64Array() []int64
	readFloat32Array() []float32
	readFloat64Array() []float64
	readStringArray() []string
}

type BinaryReaderProxy struct {
	_byteReader *bytes.Buffer
}

func newBinaryReaderProxy(data []byte) *BinaryReaderProxy {
	rs := new(BinaryReaderProxy)
	rs._byteReader = bytes.NewBuffer(data)
	return rs
}

func (o *BinaryReaderProxy) reset(data []byte) {
	o._byteReader.Reset()
	o._byteReader.Write(data)

	o._byteReader = bytes.NewBuffer(data)
}

func (o *BinaryReaderProxy) getBytesAvailable() bool {
	return o._byteReader.Len() > 0
}

//单个数据读取----------------------

func (o *BinaryReaderProxy) readBoolean() bool {
	return o.readUInt8() != 0
}

func (o *BinaryReaderProxy) readUInt8() uint8 {
	byte, _ := o._byteReader.ReadByte()
	return byte
}

func (o *BinaryReaderProxy) readUInt16() uint16 {
	return binary.BigEndian.Uint16(o._byteReader.Next(SIZE_OF_UINT16))
}

func (o *BinaryReaderProxy) readUInt32() uint32 {
	return binary.BigEndian.Uint32(o._byteReader.Next(SIZE_OF_UINT32))
}

func (o *BinaryReaderProxy) readUInt64() uint64 {
	return binary.BigEndian.Uint64(o._byteReader.Next(SIZE_OF_UINT64))
}

func (o *BinaryReaderProxy) readInt8() int8 {
	var rs int8
	binary.Read(o._byteReader, binary.BigEndian, &rs)
	return rs
}

func (o *BinaryReaderProxy) readInt16() int16 {
	var rs int16
	binary.Read(o._byteReader, binary.BigEndian, &rs)
	return rs
}

func (o *BinaryReaderProxy) readInt32() int32 {
	var rs int32
	binary.Read(o._byteReader, binary.BigEndian, &rs)
	return rs
}

func (o *BinaryReaderProxy) readInt64() int64 {
	var rs int64
	binary.Read(o._byteReader, binary.BigEndian, &rs)
	return rs
}

func (o *BinaryReaderProxy) readFloat32() float32 {
	var rs float32
	binary.Read(o._byteReader, binary.BigEndian, &rs)
	return rs
}

func (o *BinaryReaderProxy) readFloat64() float64 {
	var rs float64
	binary.Read(o._byteReader, binary.BigEndian, &rs)
	return rs
}

func (o *BinaryReaderProxy) readString() string {
	len := o.readUInt16()
	data := o._byteReader.Next(int(len))
	rs := string(data)
	return rs
}

//数组数据读取----------------------

func (o *BinaryReaderProxy) readBooleanArray() []bool {
	len := int(o.readUInt16())
	rs := make([]bool, len)
	for i := 0; i < len; i++ {
		rs[i] = o.readBoolean()
	}
	return rs
}

func (o *BinaryReaderProxy) readUInt8Array() []uint8 {
	len := int(o.readUInt16())
	rs := make([]uint8, len)
	for i := 0; i < len; i++ {
		rs[i] = o.readUInt8()
	}
	return rs
}

func (o *BinaryReaderProxy) readUInt16Array() []uint16 {
	len := int(o.readUInt16())
	rs := make([]uint16, len)
	for i := 0; i < len; i++ {
		rs[i] = o.readUInt16()
	}
	return rs
}

func (o *BinaryReaderProxy) readUInt32Array() []uint32 {
	len := int(o.readUInt16())
	rs := make([]uint32, len)
	for i := 0; i < len; i++ {
		rs[i] = o.readUInt32()
	}
	return rs
}

func (o *BinaryReaderProxy) readUInt64Array() []uint64 {
	len := int(o.readUInt16())
	rs := make([]uint64, len)
	for i := 0; i < len; i++ {
		rs[i] = o.readUInt64()
	}
	return rs
}

func (o *BinaryReaderProxy) readInt8Array() []int8 {
	len := int(o.readUInt16())
	rs := make([]int8, len)
	for i := 0; i < len; i++ {
		rs[i] = o.readInt8()
	}
	return rs
}

func (o *BinaryReaderProxy) readInt16Array() []int16 {
	len := int(o.readUInt16())
	rs := make([]int16, len)
	for i := 0; i < len; i++ {
		rs[i] = o.readInt16()
	}
	return rs
}

func (o *BinaryReaderProxy) readInt32Array() []int32 {
	len := int(o.readUInt16())
	rs := make([]int32, len)
	for i := 0; i < len; i++ {
		rs[i] = o.readInt32()
	}
	return rs
}
func (o *BinaryReaderProxy) readInt64Array() []int64 {
	len := int(o.readUInt16())
	rs := make([]int64, len)
	for i := 0; i < len; i++ {
		rs[i] = o.readInt64()
	}
	return rs
}

func (o *BinaryReaderProxy) readFloat32Array() []float32 {
	len := int(o.readUInt16())
	rs := make([]float32, len)
	for i := 0; i < len; i++ {
		rs[i] = o.readFloat32()
	}
	return rs
}

func (o *BinaryReaderProxy) readFloat64Array() []float64 {
	len := int(o.readUInt16())
	rs := make([]float64, len)
	for i := 0; i < len; i++ {
		rs[i] = o.readFloat64()
	}
	return rs
}

func (o *BinaryReaderProxy) readStringArray() []string {
	len := int(o.readUInt16())
	rs := make([]string, len)
	for i := 0; i < len; i++ {
		rs[i] = o.readString()
	}
	return rs
}
