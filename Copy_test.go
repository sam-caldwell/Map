package Map

import (
	"fmt"
	"testing"
)

func Test_Copy(t *testing.T) {
	lhs := Map[string, int]{}
	rhs := Map[string, int]{}
	for i := 0; i < 10; i++ {
		lhs[fmt.Sprintf("%d", i)] = i
	}
	Copy[string, int](lhs, rhs)
	if len(lhs) != len(rhs) {
		t.Fatal("lhs and rhs are not the same size")
	}
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("%d", i)
		if lhs[key] != rhs[key] {
			t.Fatalf("key %d mismatch", i)
		}
	}
}
