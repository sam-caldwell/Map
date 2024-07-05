package Map

// Copy - is a helper function to create a copy of the map (lhs to rhs).
func (lhs *Map[K, V]) Copy(rhs Map[K, V]) {
	rhs.Purge()
	for key, value := range *lhs {
		rhs[key] = value
	}
}
