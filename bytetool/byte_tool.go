package bytetool

import (
	"encoding/binary"
	"unsafe"
)

const (
	SHORT_SIZE uint32 = uint32(unsafe.Sizeof(uint16(0)))
	INT_SIZE   uint32 = uint32(unsafe.Sizeof(uint32(0)))
	LONG_SIZE  uint32 = uint32(unsafe.Sizeof(uint64(0)))
)

func Reverse(src []byte) []byte {
	size := len(src)
	for i := 0; i < size; i++ {
		src[i], src[size-i-1] = src[size-i-1], src[i]
		if (i + 1) == size/2 {
			break
		}
	}
	return src
}

//uint16 --> []byte
func Uint16ToBytes(n uint16) (result []byte) {
	result = make([]byte, SHORT_SIZE)
	binary.BigEndian.PutUint16(result, n)
	return
}

//[]byte --> uint16
func BytesToUint16(buf []byte) (n uint16) {
	n = binary.BigEndian.Uint16(buf)
	return
}

//uint32 --> []byte
func Uint32ToBytes(n uint32) (result []byte) {
	result = make([]byte, INT_SIZE)
	binary.BigEndian.PutUint32(result, n)
	return
}

//[]byte --> uint32
func BytesToUint32(buf []byte) (n uint32) {
	n = binary.BigEndian.Uint32(buf)
	return
}

//uint64 --> []byte
func Uint64ToBytes(n uint64) (result []byte) {
	result = make([]byte, LONG_SIZE)
	binary.BigEndian.PutUint64(result, n)
	return
}

//[]byte --> uint64
func BytesToUint64(buf []byte) (n uint64) {
	n = binary.BigEndian.Uint64(buf)
	return
}
