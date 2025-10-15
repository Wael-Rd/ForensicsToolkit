package utils

import (
	"encoding/binary"
	"unicode/utf16"
)

// UTF16leBytesToString converts UTF-16 little-endian bytes to string
func UTF16leBytesToString(leBytes []byte) string {
	if len(leBytes)%2 != 0 {
		return ""
	}

	utf16Data := make([]uint16, len(leBytes)/2)
	for i := 0; i < len(utf16Data); i++ {
		utf16Data[i] = binary.LittleEndian.Uint16(leBytes[i*2:])
	}

	return string(utf16.Decode(utf16Data))
}
