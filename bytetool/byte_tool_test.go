package bytetool

import (
	"fmt"
	"testing"
)

func Test_type(t *testing.T) {

	fmt.Println(SHORT_SIZE)
	fmt.Println(INT_SIZE)
	fmt.Println(LONG_SIZE)

	//unit16
	fmt.Println("uint16 <--> []byte")
	buf := Uint16_to_bytes(258)
	fmt.Println(buf)
	fmt.Println(Bytes_to_uint16(buf))
	r_buf := Reverse(buf)
	fmt.Println("Reverse:", r_buf)

	//unit32
	fmt.Println("uint32 <--> []byte")
	buf2 := Uint32_to_bytes(258)
	fmt.Println(buf2)
	fmt.Println(Bytes_to_uint32(buf2))
	//unit64
	fmt.Println("uint64 <--> []byte")
	buf3 := Uint64_to_bytes(258)
	fmt.Println(buf3)
	fmt.Println(Bytes_to_uint64(buf3))

}
