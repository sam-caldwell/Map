package Map

import (
	"testing"
	"unsafe"
)

func TestOverwriteData(t *testing.T) {
	// Example data
	data := []byte{1, 2, 3, 4, 5}
	ptr := unsafe.Pointer(&data[0])
	size := uintptr(len(data))

	// Call the function to overwrite the data
	overwriteData(ptr, size)

	// Check if all bytes in the slice are now zero
	for _, v := range data {
		if v != 0 {
			t.Errorf("Expected byte to be 0, got %d", v)
		}
	}
}
