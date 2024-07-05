package Map

import "testing"

func TestMap_type(t *testing.T) {
	t.Run("uninitialized state", func(t *testing.T) {
		var m Map[string, int]
		if m != nil {
			t.Fatal("expect nil")
		}
	})
	t.Run("uninitialized state", func(t *testing.T) {
		m := Map[string, int]{}
		m["one"] = 1
	})
	t.Run("initialized Map struct", func(t *testing.T) {
		m := New[string, int]()
		(*m)["one"] = 1
		(*m)["two"] = 2
	})
}
