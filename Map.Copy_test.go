package Map

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMap_Copy(t *testing.T) {
	lhs := Map[string, int]{}
	rhs := Map[string, int]{}
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("%d", i)
		lhs[key] = i
	}
	lhs.Copy(&rhs)

	if len(lhs) != len(rhs) {
		t.Fatal("copied data not equal size")
	}
	for key, actual := range lhs {
		if expected, ok := rhs[key]; !ok || !reflect.DeepEqual(actual, expected) {
			t.Fatalf("copied data values do not match\n"+
				"ok: %v\n"+
				"key: %s\n"+
				"actual: %v\n"+
				"expected: %v",
				ok, key, actual, expected)
		}
	}
}
