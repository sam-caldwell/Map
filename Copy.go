package Map

// Copy - is a helper function to create a copy of the map from lhs to rhs
func Copy[K comparable, V any](lhs, rhs map[K]V) {

	//Delete any contents from the destination to avoid merging two maps.
	Purge(&rhs)

	for key, value := range lhs {

		rhs[key] = value

	}

}
