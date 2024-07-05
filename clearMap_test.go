package Map

import (
	"reflect"
	"testing"
)

func TestClearMap(t *testing.T) {
	// Create a map with some initial data
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	// Ensure the map has elements before clearing
	if len(m) != 3 {
		t.Errorf("Expected map length 3, got %d", len(m))
	}

	// Clear the map
	clearMap(m)

	// Check if the map is empty after clearing
	if len(m) != 0 {
		t.Errorf("Expected map length 0 after clearing, got %d", len(m))
	}

	// Additionally, you can verify that the map type remains the same
	expectedType := reflect.TypeOf(map[string]int{})
	actualType := reflect.TypeOf(m)
	if expectedType != actualType {
		t.Errorf("Expected map type %s, got %s", expectedType, actualType)
	}
}
