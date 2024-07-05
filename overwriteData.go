package Map

import (
	"unsafe"
)

// overwriteData overwrites the contents of a byte slice with zeros
func overwriteData(ptr unsafe.Pointer, size uintptr) {
	dataBytes := (*[1 << 30]byte)(ptr)[:size]
	for i := range dataBytes {
		dataBytes[i] = 0x00
	}
}
