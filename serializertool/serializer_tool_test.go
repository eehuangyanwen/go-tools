package serializertool

import (
	"fmt"
	"testing"
)

var p = fmt.Println

func Test_type(t *testing.T) {
	b1 := byte(1)
	b2 := uint16(2)
	b3 := uint32(3)
	b4 := uint64(4)
	str := "我勒个去啊，Let's go"

	//
	var wb WriteBuffer

	p(wb)
	wb.WriteByte(b1)
	wb.WriteUint16(b2)
	wb.WriteUint32(b3)
	wb.WriteUint64(b4)
	wb.WriteString(str)

	buf := wb.Bytes()
	p(buf)

	//
	rb := NewReadBuffer(buf)

	_b1, _ := rb.ReadByte()
	p(_b1)

	_b2, _ := rb.ReadUint16()
	p(_b2)

	_b3, _ := rb.ReadUint32()
	p(_b3)

	_b4, _ := rb.ReadUint64()
	p(_b4)

	_str, _ := rb.ReadString()
	p(_str)
}
