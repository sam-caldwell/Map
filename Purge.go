package Map

import "runtime"

// Purge - reset the memory state of a given generic map and call gc explicitly.
func Purge[K comparable, V any](mp *map[K]V) {
	m := make(map[K]V, 1)
	mp = &m
	runtime.GC()
}
