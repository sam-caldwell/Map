package Map

import (
	"testing"
)

func TestMap_Equal(t *testing.T) {
	t.Run("lhs nil, rhs is not", func(t *testing.T) {
		var lhs Map[string, int]
		if lhs != nil {
			t.Fatal("lhs should be nil")
		}
		rhs := Map[string, int]{}
		rhs["1"] = 1
		if lhs.Equal(&rhs) {
			t.Fatal("expected inequality")
		}
	})
	t.Run("lhs is not nil, rhs is nil", func(t *testing.T) {
		var rhs Map[string, int]
		if rhs != nil {
			t.Fatal("rhs should be nil")
		}
		lhs := Map[string, int]{}
		lhs["1"] = 1
		if lhs.Equal(&rhs) {
			t.Fatal("expected inequality")
		}
	})
	t.Run("lhs and rhs not equal (size)", func(t *testing.T) {
		lhs := Map[string, int]{
			"1": 1,
			"2": 2,
			"3": 3,
		}
		rhs := Map[string, int]{
			"1": 1,
			"2": 2,
		}
		if lhs.Equal(&rhs) {
			t.Fatal("lhs != rhs")
		}
		if rhs.Equal(&lhs) {
			t.Fatal("rhs != lhs")
		}
	})
	t.Run("lhs and rhs not equal keys", func(t *testing.T) {
		lhs := Map[string, int]{
			"1": 1,
			"2": 2,
			"3": 3,
		}
		rhs := Map[string, int]{
			"1":   1,
			"2":   2,
			"bad": 3,
		}
		if lhs.Equal(&rhs) {
			t.Fatal("lhs != rhs")
		}
		if rhs.Equal(&lhs) {
			t.Fatal("rhs != lhs")
		}
	})
	t.Run("lhs and rhs not equal values", func(t *testing.T) {
		lhs := Map[string, int]{
			"1": 1,
			"2": 2,
			"3": 3,
		}
		rhs := Map[string, int]{
			"1": 1,
			"2": 2,
			"3": 6,
		}
		if lhs.Equal(&rhs) {
			t.Fatal("lhs != rhs")
		}
		if rhs.Equal(&lhs) {
			t.Fatal("rhs != lhs")
		}
	})
}
