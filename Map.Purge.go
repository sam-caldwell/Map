package Map

import (
	"runtime"
)

// Purge - reset the memory state of a given generic map and call gc explicitly.
func (lhs *Map[K, V]) Purge() {
	for k, _ := range (*lhs)[K, V] {
		sanitizeData((*lhs)[k])
	}
	*lhs = make(Map[K, V], 1)
	runtime.GC()
}
