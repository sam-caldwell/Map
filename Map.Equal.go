package Map

import "reflect"

// Equal - Compare two maps and return boolean indicating equivalence
func (lhs *Map[K, V]) Equal(rhs *Map[K, V]) bool {
	if lhs == nil || rhs == nil {
		return lhs == nil && rhs == nil
	}
	if len(*lhs) != len(*rhs) {
		return false
	}
	for key, lValue := range *lhs {
		if rValue, ok := (*rhs)[key]; !ok || !reflect.DeepEqual(lValue, rValue) {
			return false
		}
	}
	return true
}
