package Map

import (
	"fmt"
	"testing"
)

func TestEqual(t *testing.T) {
	t.Run("equal array", func(t *testing.T) {
		lhs := map[string]int{}
		rhs := map[string]int{}
		t.Run("setup", func(t *testing.T) {
			for i := 0; i < 10; i++ {
				lhs[fmt.Sprintf("%d", i)] = i
				rhs[fmt.Sprintf("%d", i)] = i
			}
		})
		t.Run("verify", func(t *testing.T) {
			if !Equal[string, int](&lhs, &rhs) {
				t.Fatal("arrays should be the same")
			}
		})
	})
	t.Run("nil lhs", func(t *testing.T) {
		lhs := make(map[string]int)
		lhs["1"] = 1
		rhs := make(map[string]int)
		rhs["1"] = 1
		if Equal[string, int](nil, &rhs) {
			t.Fatalf("expect lhs nil, rhs not nil")
		}
		if !Equal[string, int](nil, nil) {
			t.Fatalf("expect lhs nil, rhs nil")
		}
		if !Equal[string, int](&lhs, &rhs) {
			t.Fatalf("expect lhs and rhs are equal")
		}
	})
	t.Run("inequality (size)", func(t *testing.T) {
		lhs := make(map[string]int)
		lhs["1"] = 1
		lhs["2"] = 2
		rhs := make(map[string]int)
		rhs["1"] = 1
		if Equal[string, int](&lhs, &rhs) {
			t.Fatalf("expect lhs and rhs are not equal (size)")
		}
	})
	t.Run("inequality (differing values)", func(t *testing.T) {
		t.Run("expect lhs and rhs are not equal values(1)", func(t *testing.T) {
			lhs := make(map[string]int)
			lhs["1"] = 1
			lhs["2"] = 2
			rhs := make(map[string]int)
			rhs["1"] = 1
			rhs["2"] = 3
			if Equal[string, int](&lhs, &rhs) {
				t.Fail()
			}
		})
		t.Run("expect lhs and rhs are not equal values(2)", func(t *testing.T) {
			lhs := make(map[string]int)
			lhs["1"] = 1
			lhs["2"] = 3
			rhs := make(map[string]int)
			rhs["1"] = 1
			rhs["2"] = 2
			if Equal[string, int](&lhs, &rhs) {
				t.Fail()
			}
		})
		t.Run("expect lhs and rhs are not equal values(3)", func(t *testing.T) {
			lhs := make(map[string]int)
			lhs["1"] = 1
			lhs["3"] = 3
			rhs := make(map[string]int)
			rhs["1"] = 1
			rhs["2"] = 2
			if Equal[string, int](&lhs, &rhs) {
				t.Fail()
			}
		})
	})
}
