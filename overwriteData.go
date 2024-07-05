package Map

import (
	"unsafe"
)

// clearMap clears the contents of a map
func clearMap[K comparable, V any](m map[K]V) {
	// Iterate through the map and set each key-value pair to the zero value
	for k := range m {
		delete(m, k)
	}
}

// overwriteData overwrites the contents of a byte slice with zeros
func overwriteData(ptr unsafe.Pointer, size uintptr) {
	dataBytes := (*[1 << 30]byte)(ptr)[:size]
	for i := range dataBytes {
		dataBytes[i] = 0x00
	}
}
