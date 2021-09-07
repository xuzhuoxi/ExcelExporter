package proxy2

import (
	"bytes"
	"encoding/binary"
)

const (
	SIZE_OF_BOOLEAN = 1
	SIZE_OF_INT8    = 1
	SIZE_OF_INT16   = 2
	SIZE_OF_INT32   = 4
	SIZE_OF_INT64   = 8
	SIZE_OF_UINT8   = 1
	SIZE_OF_UINT16  = 2
	SIZE_OF_UINT32  = 4
	SIZE_OF_UINT64  = 8
	SIZE_OF_FLOAT32 = 4
	SIZE_OF_FLOAT64 = 8
)

type IBinaryReaderProxy interface {
	//----------------------
	SetData(data []byte)
	GetBytesAvailable() bool
	//单个数据读取----------------------
	ReadBoolean() bool
	ReadUInt8() uint8
	ReadUInt16() uint16
	ReadUInt32() uint32
	ReadUInt64() uint64
	ReadInt8() int8
	ReadInt16() int16
	ReadInt32() int32
	ReadInt64() int64
	ReadFloat32() float32
	ReadFloat64() float64
	ReadString() string
	//数组数据读取----------------------
	ReadBooleanArray() []bool
	ReadUInt8Array() []uint8
	ReadUInt16Array() []uint16
	ReadUInt32Array() []uint32
	ReadUInt64Array() []uint64
	ReadInt8Array() []int8
	ReadInt16Array() []int16
	ReadInt32Array() []int32
	ReadInt64Array() []int64
	ReadFloat32Array() []float32
	ReadFloat64Array() []float64
	ReadStringArray() []string
}

func NewBinaryReaderProxy(data []byte) *BinaryReaderProxy {
	rs := new(BinaryReaderProxy)
	rs.ByteReader = bytes.NewBuffer(data)
	return rs
}

type BinaryReaderProxy struct {
	ByteReader *bytes.Buffer
}

func (o *BinaryReaderProxy) SetData(data []byte) {
	o.ByteReader.Reset()
	o.ByteReader.Write(data)

	o.ByteReader = bytes.NewBuffer(data)
}

func (o *BinaryReaderProxy) GetBytesAvailable() bool {
	return o.ByteReader.Len() > 0
}

//单个数据读取----------------------

func (o *BinaryReaderProxy) ReadBoolean() bool {
	return o.ReadUInt8() != 0
}

func (o *BinaryReaderProxy) ReadUInt8() uint8 {
	byte, _ := o.ByteReader.ReadByte()
	return byte
}

func (o *BinaryReaderProxy) ReadUInt16() uint16 {
	return binary.BigEndian.Uint16(o.ByteReader.Next(SIZE_OF_UINT16))
}

func (o *BinaryReaderProxy) ReadUInt32() uint32 {
	return binary.BigEndian.Uint32(o.ByteReader.Next(SIZE_OF_UINT32))
}

func (o *BinaryReaderProxy) ReadUInt64() uint64 {
	return binary.BigEndian.Uint64(o.ByteReader.Next(SIZE_OF_UINT64))
}

func (o *BinaryReaderProxy) ReadInt8() int8 {
	var rs int8
	binary.Read(o.ByteReader, binary.BigEndian, &rs)
	return rs
}

func (o *BinaryReaderProxy) ReadInt16() int16 {
	var rs int16
	binary.Read(o.ByteReader, binary.BigEndian, &rs)
	return rs
}

func (o *BinaryReaderProxy) ReadInt32() int32 {
	var rs int32
	binary.Read(o.ByteReader, binary.BigEndian, &rs)
	return rs
}

func (o *BinaryReaderProxy) ReadInt64() int64 {
	var rs int64
	binary.Read(o.ByteReader, binary.BigEndian, &rs)
	return rs
}

func (o *BinaryReaderProxy) ReadFloat32() float32 {
	var rs float32
	binary.Read(o.ByteReader, binary.BigEndian, &rs)
	return rs
}

func (o *BinaryReaderProxy) ReadFloat64() float64 {
	var rs float64
	binary.Read(o.ByteReader, binary.BigEndian, &rs)
	return rs
}

func (o *BinaryReaderProxy) ReadString() string {
	len := o.ReadUInt16()
	data := o.ByteReader.Next(int(len))
	rs := string(data)
	return rs
}

//数组数据读取----------------------

func (o *BinaryReaderProxy) ReadBooleanArray() []bool {
	len := int(o.ReadUInt16())
	rs := make([]bool, len)
	for i := 0; i < len; i++ {
		rs[i] = o.ReadBoolean()
	}
	return rs
}

func (o *BinaryReaderProxy) ReadUInt8Array() []uint8 {
	len := int(o.ReadUInt16())
	rs := make([]uint8, len)
	for i := 0; i < len; i++ {
		rs[i] = o.ReadUInt8()
	}
	return rs
}

func (o *BinaryReaderProxy) ReadUInt16Array() []uint16 {
	len := int(o.ReadUInt16())
	rs := make([]uint16, len)
	for i := 0; i < len; i++ {
		rs[i] = o.ReadUInt16()
	}
	return rs
}

func (o *BinaryReaderProxy) ReadUInt32Array() []uint32 {
	len := int(o.ReadUInt16())
	rs := make([]uint32, len)
	for i := 0; i < len; i++ {
		rs[i] = o.ReadUInt32()
	}
	return rs
}

func (o *BinaryReaderProxy) ReadUInt64Array() []uint64 {
	len := int(o.ReadUInt16())
	rs := make([]uint64, len)
	for i := 0; i < len; i++ {
		rs[i] = o.ReadUInt64()
	}
	return rs
}

func (o *BinaryReaderProxy) ReadInt8Array() []int8 {
	len := int(o.ReadUInt16())
	rs := make([]int8, len)
	for i := 0; i < len; i++ {
		rs[i] = o.ReadInt8()
	}
	return rs
}

func (o *BinaryReaderProxy) ReadInt16Array() []int16 {
	len := int(o.ReadUInt16())
	rs := make([]int16, len)
	for i := 0; i < len; i++ {
		rs[i] = o.ReadInt16()
	}
	return rs
}

func (o *BinaryReaderProxy) ReadInt32Array() []int32 {
	len := int(o.ReadUInt16())
	rs := make([]int32, len)
	for i := 0; i < len; i++ {
		rs[i] = o.ReadInt32()
	}
	return rs
}
func (o *BinaryReaderProxy) ReadInt64Array() []int64 {
	len := int(o.ReadUInt16())
	rs := make([]int64, len)
	for i := 0; i < len; i++ {
		rs[i] = o.ReadInt64()
	}
	return rs
}

func (o *BinaryReaderProxy) ReadFloat32Array() []float32 {
	len := int(o.ReadUInt16())
	rs := make([]float32, len)
	for i := 0; i < len; i++ {
		rs[i] = o.ReadFloat32()
	}
	return rs
}

func (o *BinaryReaderProxy) ReadFloat64Array() []float64 {
	len := int(o.ReadUInt16())
	rs := make([]float64, len)
	for i := 0; i < len; i++ {
		rs[i] = o.ReadFloat64()
	}
	return rs
}

func (o *BinaryReaderProxy) ReadStringArray() []string {
	len := int(o.ReadUInt16())
	rs := make([]string, len)
	for i := 0; i < len; i++ {
		rs[i] = o.ReadString()
	}
	return rs
}
