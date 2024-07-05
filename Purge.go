package Map

import (
	"reflect"
	"runtime"
	"unsafe"
)

// Purge - reset the memory state of a given generic map and call gc explicitly.
func Purge[K comparable, V any](mp *map[K]V) {
	if mp == nil {
		return
	}

	// Use reflect to get the pointer and size of the map
	mpValue := reflect.ValueOf(*mp)
	ptr := unsafe.Pointer(mpValue.Pointer())
	size := uintptr(mpValue.Type().Size())

	// Overwrite the data in the map
	overwriteData(ptr, size)

	// Clear the map and reset it
	clearMap(*mp)
	*mp = make(map[K]V) // Create a new empty map of the same type

	runtime.GC()
}
