package Map

// Copy - is a helper function to create a copy of the map.
func Copy[K comparable, V any](dest, src map[K]V) {
	Purge(&dest) //Delete any contents from the destination to avoid merging two maps.
	for key, value := range src {
		dest[key] = value
	}
}
