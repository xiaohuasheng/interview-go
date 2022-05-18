package gopractice

import (
	"fmt"
)

func endian() {
	s := int16(0x1234)
	b := int8(s)
	if 0x34 == b {
		fmt.Println("little endian")
	} else {
		fmt.Println("big endian")
	}
}
