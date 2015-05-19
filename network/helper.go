package network

import (
	"bytes"
	"encoding/binary"
)

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, x)
	return bytesBuffer.Bytes()
}

//浮点数转换成字节
func FloatToBytes(n float32) []byte {
	x := float32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, x)
	return bytesBuffer.Bytes()
}

//双精度转换成字节
func DoubleToBytes(n float64) []byte {
	x := float64(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.LittleEndian, &x)
	return int(x)
}

//字节转换成浮点数
func BytesToFloat(b []byte) float32 {
	bytesBuffer := bytes.NewBuffer(b)
	var x float32
	binary.Read(bytesBuffer, binary.LittleEndian, &x)
	return float32(x)
}

//字节转换成双精度
func BytesToDouble(b []byte) float64 {
	bytesBuffer := bytes.NewBuffer(b)
	var x float64
	binary.Read(bytesBuffer, binary.LittleEndian, &x)
	return float64(x)
}

//字节转字符串
func BytesToString(b []byte) string {
	bytesBuffer := bytes.NewBuffer(b)
	var s string
	binary.Read(bytesBuffer, binary.LittleEndian, &s)
	return string(s)
}
