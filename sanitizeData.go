package Map

import (
	"reflect"
	"unsafe"
)

// sanitizeData - Sanitize the binary representation of a value
func sanitizeData(value interface{}) {
	// Use unsafe.Pointer to get the address of the value's memory representation
	ptr := unsafe.Pointer(reflect.ValueOf(value).Pointer())

	// Determine the size of the value in memory
	size := unsafe.Sizeof(value)

	// Convert the pointer to a byte slice
	dataBytes := (*[1 << 30]byte)(ptr)[:size]

	// Set the byte slice to nil
	for i := range dataBytes {
		dataBytes[i] = 0x00
	}
}
