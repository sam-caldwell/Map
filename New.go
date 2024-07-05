package Map

// New - Initialize and return a generic map instance
func New[K comparable, V any]() *Map[K, V] {

	m := make(Map[K, V], 1)

	return &m

}
