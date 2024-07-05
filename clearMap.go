package Map

// clearMap clears the contents of a map
func clearMap[K comparable, V any](m map[K]V) {
	// Iterate through the map and set each key-value pair to the zero value
	for k := range m {
		delete(m, k)
	}
}
